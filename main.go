/*
 * Nucmf_UECapabilityManagement
 *
 * Nucmf_UECapabilityManagement Service. © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
)

func main() {
	log.Printf("Server started")

	ADictionaryEntryDocumentApiService := openapi.NewADictionaryEntryDocumentApiService()
	ADictionaryEntryDocumentApiController := openapi.NewADictionaryEntryDocumentApiController(ADictionaryEntryDocumentApiService)

	DicEntryIdDocumentApiService := openapi.NewDicEntryIdDocumentApiService()
	DicEntryIdDocumentApiController := openapi.NewDicEntryIdDocumentApiController(DicEntryIdDocumentApiService)

	DictionaryEntryStoreApiService := openapi.NewDictionaryEntryStoreApiService()
	DictionaryEntryStoreApiController := openapi.NewDictionaryEntryStoreApiController(DictionaryEntryStoreApiService)

	IndividualSubscriptionDocumentApiService := openapi.NewIndividualSubscriptionDocumentApiService()
	IndividualSubscriptionDocumentApiController := openapi.NewIndividualSubscriptionDocumentApiController(IndividualSubscriptionDocumentApiService)

	SubscriptionsCollectionApiService := openapi.NewSubscriptionsCollectionApiService()
	SubscriptionsCollectionApiController := openapi.NewSubscriptionsCollectionApiController(SubscriptionsCollectionApiService)

	router := openapi.NewRouter(ADictionaryEntryDocumentApiController, DicEntryIdDocumentApiController, DictionaryEntryStoreApiController, IndividualSubscriptionDocumentApiController, SubscriptionsCollectionApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}