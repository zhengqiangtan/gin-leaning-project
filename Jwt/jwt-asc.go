package main

import (
	"context"
	"fmt"
	"github.com/cidertool/asc-go/asc"
	"os"
	"time"
)

func main() {
	// Key ID for the given private key, described in App Store Connect
	keyID := "T22Y827GNG"
	// Issuer ID for the App Store Connect team
	issuerID := "10696be8-e055-47c7-af84-d934999b7abf"
	// A duration value for the lifetime of a token. App Store Connect does not accept a token with a lifetime of longer than 20 minutes
	expiryDuration := 20 * time.Minute
	// The bytes of the PKCS#8 private key created on App Store Connect. Keep this key safe as you can only download it once.
	privateKey, _ := os.ReadFile("C:\\Users\\HNK7WC3\\Downloads\\AuthKey_T22Y827GNG.p8")

	auth, err := asc.NewTokenConfig(keyID, issuerID, expiryDuration, privateKey)

	if err != nil {
		print(err.Error())
	}
	client := asc.NewClient(auth.Client())

	//// list all apps with the bundle ID "com.sky.MyApp" in the authenticated user's team
	//apps, _, err := client.Apps.ListApps(&asc.ListAppsQuery{
	//	FilterBundleID: []string{"com.sky.MyApp"},
	//})
	//client.Builds.

	fmt.Println(client.TestFlight, client.UserAgent)
	ctx := context.Background()

	//client.Submission.GetReviewDetail(ctx, "1468073139", nil)

	client.Apps.ListAppInfosForApp(ctx, "1468073139", nil)

	//
	//fmt.Println(response,i)

}
