package db

import (
	"reflect"
	"testing"
	"time"

	"github.com/lib/pq"
	"github.com/sourcegraph/sourcegraph/pkg/errcode"
)

// registryExtensionNamesForTests is a list of test cases containing valid and invalid registry
// extension names.
var registryExtensionNamesForTests = []struct {
	name      string
	wantValid bool
}{
	{"", false},
	{"a", true},
	{"-a", false},
	{"a-", false},
	{"a-b", true},
	{"a--b", false},
	{"a---b", false},
	{"a.b", true},
	{"a..b", false},
	{"a...b", false},
	{"a_b", true},
	{"a__b", false},
	{"a___b", false},
	{"a-.b", false},
	{"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", false},
}

func TestRegistryExtensions_validUsernames(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := testContext()

	user, err := Users.Create(ctx, NewUser{Username: "u"})
	if err != nil {
		t.Fatal(err)
	}

	for _, test := range registryExtensionNamesForTests {
		t.Run(test.name, func(t *testing.T) {
			valid := true
			if _, err := RegistryExtensions.Create(ctx, user.ID, 0, test.name); err != nil {
				if e, ok := err.(*pq.Error); ok && (e.Constraint == "registry_extensions_name_valid_chars" || e.Constraint == "registry_extensions_name_length") {
					valid = false
				} else {
					t.Fatal(err)
				}
			}
			if valid != test.wantValid {
				t.Errorf("%q: got valid %v, want %v", test.name, valid, test.wantValid)
			}
		})
	}
}

func TestRegistryExtensions(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := testContext()

	testGetByID := func(t *testing.T, id int32, want *RegistryExtension, wantPublisherName string) {
		t.Helper()
		x, err := RegistryExtensions.GetByID(ctx, id)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(x, want) {
			t.Errorf("got %+v, want %+v", x, want)
		}
		if x.Publisher.NonCanonicalName != wantPublisherName {
			t.Errorf("got publisher name %q, want %q", x.Publisher.NonCanonicalName, wantPublisherName)
		}
	}
	testGetByExtensionID := func(t *testing.T, extensionID string, want *RegistryExtension) {
		t.Helper()
		x, err := RegistryExtensions.GetByExtensionID(ctx, extensionID)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(x, want) {
			t.Errorf("got %+v, want %+v", x, want)
		}
		if x.NonCanonicalExtensionID != extensionID {
			t.Errorf("got extension ID %q, want %q", x.NonCanonicalExtensionID, extensionID)
		}
	}
	testList := func(t *testing.T, opt RegistryExtensionsListOptions, want []*RegistryExtension) {
		t.Helper()
		if ois, err := RegistryExtensions.List(ctx, opt); err != nil {
			t.Fatal(err)
		} else if !reflect.DeepEqual(ois, want) {
			t.Errorf("got %s, want %s", asJSON(t, ois), asJSON(t, want))
		}
	}
	testListCount := func(t *testing.T, opt RegistryExtensionsListOptions, want []*RegistryExtension) {
		t.Helper()
		testList(t, opt, want)
		if n, err := RegistryExtensions.Count(ctx, opt); err != nil {
			t.Fatal(err)
		} else if want := len(want); n != want {
			t.Errorf("got %d, want %d", n, want)
		}
	}

	user, err := Users.Create(ctx, NewUser{Username: "u"})
	if err != nil {
		t.Fatal(err)
	}
	org, err := Orgs.Create(ctx, "o", nil)
	if err != nil {
		t.Fatal(err)
	}

	createAndGet := func(t *testing.T, publisherUserID, publisherOrgID int32, name string) *RegistryExtension {
		t.Helper()
		xID, err := RegistryExtensions.Create(ctx, publisherUserID, publisherOrgID, name)
		if err != nil {
			t.Fatal(err)
		}
		x, err := RegistryExtensions.GetByID(ctx, xID)
		if err != nil {
			t.Fatal(err)
		}
		return x
	}
	xu := createAndGet(t, user.ID, 0, "xu")
	xo := createAndGet(t, 0, org.ID, "xo")

	t.Run("List/Count publishers", func(t *testing.T) {
		publishers, err := RegistryExtensions.ListPublishers(ctx, RegistryPublishersListOptions{})
		if err != nil {
			t.Fatal(err)
		}
		if want := []*RegistryPublisher{
			&xo.Publisher,
			&xu.Publisher,
		}; !reflect.DeepEqual(publishers, want) {
			t.Errorf("got publishers %+v, want %+v", publishers, want)
		}

		if n, err := RegistryExtensions.CountPublishers(ctx, RegistryPublishersListOptions{}); err != nil {
			t.Fatal(err)
		} else if want := 2; n != 2 {
			t.Errorf("got count %d, want %d", n, want)
		}
	})

	publishers := map[string]struct {
		publisherUserID, publisherOrgID int32
		publisherName                   string
	}{
		"user": {publisherUserID: user.ID, publisherName: "u"},
		"org":  {publisherOrgID: org.ID, publisherName: "o"},
	}
	for name, c := range publishers {
		t.Run(name+" publisher", func(t *testing.T) {
			x := createAndGet(t, c.publisherUserID, c.publisherOrgID, "x")

			t.Run("GetByID", func(t *testing.T) {
				testGetByID(t, x.ID, x, c.publisherName)
				if _, err := RegistryExtensions.GetByID(ctx, 12345 /* doesn't exist */); !errcode.IsNotFound(err) {
					t.Errorf("got err %v, want errcode.IsNotFound", err)
				}
			})

			t.Run("GetByExtensionID", func(t *testing.T) {
				testGetByExtensionID(t, c.publisherName+"/"+x.Name, x)
				if _, err := RegistryExtensions.GetByExtensionID(ctx, "foo.bar"); !errcode.IsNotFound(err) {
					t.Errorf("got err %v, want errcode.IsNotFound", err)
				}
			})

			t.Run("List/Count all", func(t *testing.T) {
				testListCount(t, RegistryExtensionsListOptions{}, []*RegistryExtension{xu, xo, x})
			})
			wantByPublisherUser := []*RegistryExtension{xu}
			wantByPublisherOrg := []*RegistryExtension{xo}
			var wantByCurrent []*RegistryExtension
			if c.publisherUserID != 0 {
				wantByPublisherUser = append(wantByPublisherUser, x)
				wantByCurrent = wantByPublisherUser
			} else {
				wantByPublisherOrg = append(wantByPublisherOrg, x)
				wantByCurrent = wantByPublisherOrg
			}
			t.Run("List/Count by PublisherUserID", func(t *testing.T) {
				testListCount(t, RegistryExtensionsListOptions{Publisher: RegistryPublisher{UserID: user.ID}}, wantByPublisherUser)
			})
			t.Run("List/Count by Publisher.OrgID", func(t *testing.T) {
				testListCount(t, RegistryExtensionsListOptions{Publisher: RegistryPublisher{OrgID: org.ID}}, wantByPublisherOrg)
			})
			t.Run("List/Count by Publisher.Query all", func(t *testing.T) {
				testListCount(t, RegistryExtensionsListOptions{Query: "x"}, []*RegistryExtension{xu, xo, x})
			})
			t.Run("List/Count by Publisher.Query one", func(t *testing.T) {
				testListCount(t, RegistryExtensionsListOptions{Query: c.publisherName + "/" + x.Name}, wantByCurrent)
			})
			t.Run("List/Count with prioritizeExtensionIDs", func(t *testing.T) {
				testList(t, RegistryExtensionsListOptions{PrioritizeExtensionIDs: []string{xu.NonCanonicalExtensionID}, LimitOffset: &LimitOffset{Limit: 1}}, []*RegistryExtension{xu})
				testList(t, RegistryExtensionsListOptions{PrioritizeExtensionIDs: []string{xo.NonCanonicalExtensionID}, LimitOffset: &LimitOffset{Limit: 1}}, []*RegistryExtension{xo})
			})

			if err := RegistryExtensions.Delete(ctx, x.ID); err != nil {
				t.Fatal(err)
			}
			if err := RegistryExtensions.Delete(ctx, x.ID); !errcode.IsNotFound(err) {
				t.Errorf("2nd Delete: got err %v, want errcode.IsNotFound", err)
			}
			if _, err := RegistryExtensions.GetByID(ctx, x.ID); !errcode.IsNotFound(err) {
				t.Errorf("GetByID after Delete: got err %v, want errcode.IsNotFound", err)
			}
		})
	}

	t.Run("Update", func(t *testing.T) {
		x := xu
		const manifest = `"d"`
		if err := RegistryExtensions.Update(ctx, x.ID, nil, strptr(manifest)); err != nil {
			t.Fatal(err)
		}
		x1, err := RegistryExtensions.GetByID(ctx, x.ID)
		if err != nil {
			t.Fatal(err)
		}
		if time.Since(x1.UpdatedAt) > 1*time.Minute {
			t.Errorf("got UpdatedAt %v, want recent", x1.UpdatedAt)
		}
		if x1.Name != x.Name {
			t.Errorf("got name %q, want %q", *x1.Manifest, x.Name)
		}
		if x1.Manifest == nil || (*x1.Manifest != manifest) {
			t.Errorf("got data %q, want %q", *x1.Manifest, manifest)
		}
	})
}