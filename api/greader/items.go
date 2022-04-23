package greader

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/kwo/stringer/models"
)

type ItemListProvider interface {
	GetItemsUnread(max int) (models.Items, error)
}

type ItemReferenceList struct {
	ItemReferences []ItemReference `json:"itemRefs"`
}

type ItemReference struct {
	ID        string   `json:"id"`
	StreamIDs []string `json:"directStreamIds"`
	Timestamp int64    `json:"timestampUsec"`
}

func itemsIDs(itemListProvider ItemListProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// n            10000
		// s streamID   user/-/state/com.google/reading-list
		// xt exclude   user/-/state/com.google/read
		// output       json

		/*
			itemRefs: [
				{
				  id: "3614359203",
				  directStreamIds: [
					"user/1005921515/label/MTB"
				  ],
			      timestampUsec: "1416313130505268" nanoseconds
				},
		*/

		ctx := r.Context()
		n := chi.URLParamFromCtx(ctx, "n")

		maxCount := 1000
		if x, err := strconv.Atoi(n); err == nil {
			maxCount = x
		} else {
			log.Printf("invalid number: %s", err)
		}

		if user := getUserFromContext(ctx); user != nil {

			items, err := itemListProvider.GetItemsUnread(maxCount)
			if err != nil {
				log.Printf("cannot retrieve items: %s", err)
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			itemReferenceList := &ItemReferenceList{
				ItemReferences: makeItemReferences(items),
			}

			data, err := json.Marshal(itemReferenceList)
			if err != nil {
				log.Printf("item reference list: cannot marshal response: %s", err)
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			w.Header().Set(hContentType, mimetypeJson)
			_, _ = w.Write(data)
			return

		}

		http.Error(w, "", http.StatusUnauthorized)

	}
}

func makeItemReferences(items models.Items) []ItemReference {

	var result []ItemReference

	for _, item := range items {

		result = append(result, ItemReference{
			ID:        item.ID,
			StreamIDs: []string{},
			// Timestamp: item.,
		})

	}

	return result

}
