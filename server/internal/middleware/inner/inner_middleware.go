// GENERATED CODE - DO NOT EDIT!
// @generated
//
// Generated by:
//
//   go run gen_middleware.go
//
// Called via:
//
//   go generate
//

package inner

import (
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"sourcegraph.com/sourcegraph/sourcegraph/go-sourcegraph/sourcegraph"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/inventory"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/vcs"
	"sourcegraph.com/sourcegraph/sourcegraph/server/internal/middleware/inner/trace"
	"sourcegraph.com/sourcegraph/sourcegraph/server/local"
	"sourcegraph.com/sourcegraph/sourcegraph/services/svc"
	"sourcegraph.com/sourcegraph/srclib/store/pb"
	"sourcegraph.com/sqs/pbtypes"
)

// Services returns the local services wrapped with auth, etc.
func Services() svc.Services {
	return svc.Services{

		MultiRepoImporter: wrappedMultiRepoImporter{},

		Accounts: wrappedAccounts{},

		Annotations: wrappedAnnotations{},

		Auth: wrappedAuth{},

		Builds: wrappedBuilds{},

		Defs: wrappedDefs{},

		Deltas: wrappedDeltas{},

		Meta: wrappedMeta{},

		MirrorRepos: wrappedMirrorRepos{},

		Notify: wrappedNotify{},

		Orgs: wrappedOrgs{},

		People: wrappedPeople{},

		RepoStatuses: wrappedRepoStatuses{},

		RepoTree: wrappedRepoTree{},

		Repos: wrappedRepos{},

		Search: wrappedSearch{},

		Users: wrappedUsers{},
	}
}

type wrappedMultiRepoImporter struct{}

func (s wrappedMultiRepoImporter) Import(ctx context.Context, param *pb.ImportOp) (res *pbtypes.Void, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "MultiRepoImporter", "Import", param)
	defer func() {
		trace.After(ctx, "MultiRepoImporter", "Import", param, err, time.Since(start))
	}()
	res, err = local.Services.MultiRepoImporter.Import(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "MultiRepoImporter.Import returned nil, nil")
	}
	return
}

func (s wrappedMultiRepoImporter) CreateVersion(ctx context.Context, param *pb.CreateVersionOp) (res *pbtypes.Void, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "MultiRepoImporter", "CreateVersion", param)
	defer func() {
		trace.After(ctx, "MultiRepoImporter", "CreateVersion", param, err, time.Since(start))
	}()
	res, err = local.Services.MultiRepoImporter.CreateVersion(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "MultiRepoImporter.CreateVersion returned nil, nil")
	}
	return
}

func (s wrappedMultiRepoImporter) Index(ctx context.Context, param *pb.IndexOp) (res *pbtypes.Void, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "MultiRepoImporter", "Index", param)
	defer func() {
		trace.After(ctx, "MultiRepoImporter", "Index", param, err, time.Since(start))
	}()
	res, err = local.Services.MultiRepoImporter.Index(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "MultiRepoImporter.Index returned nil, nil")
	}
	return
}

type wrappedAccounts struct{}

func (s wrappedAccounts) Create(ctx context.Context, param *sourcegraph.NewAccount) (res *sourcegraph.UserSpec, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Accounts", "Create", param)
	defer func() {
		trace.After(ctx, "Accounts", "Create", param, err, time.Since(start))
	}()
	res, err = local.Services.Accounts.Create(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Accounts.Create returned nil, nil")
	}
	return
}

func (s wrappedAccounts) RequestPasswordReset(ctx context.Context, param *sourcegraph.PersonSpec) (res *sourcegraph.PendingPasswordReset, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Accounts", "RequestPasswordReset", param)
	defer func() {
		trace.After(ctx, "Accounts", "RequestPasswordReset", param, err, time.Since(start))
	}()
	res, err = local.Services.Accounts.RequestPasswordReset(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Accounts.RequestPasswordReset returned nil, nil")
	}
	return
}

func (s wrappedAccounts) ResetPassword(ctx context.Context, param *sourcegraph.NewPassword) (res *pbtypes.Void, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Accounts", "ResetPassword", param)
	defer func() {
		trace.After(ctx, "Accounts", "ResetPassword", param, err, time.Since(start))
	}()
	res, err = local.Services.Accounts.ResetPassword(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Accounts.ResetPassword returned nil, nil")
	}
	return
}

func (s wrappedAccounts) Update(ctx context.Context, param *sourcegraph.User) (res *pbtypes.Void, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Accounts", "Update", param)
	defer func() {
		trace.After(ctx, "Accounts", "Update", param, err, time.Since(start))
	}()
	res, err = local.Services.Accounts.Update(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Accounts.Update returned nil, nil")
	}
	return
}

func (s wrappedAccounts) Delete(ctx context.Context, param *sourcegraph.PersonSpec) (res *pbtypes.Void, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Accounts", "Delete", param)
	defer func() {
		trace.After(ctx, "Accounts", "Delete", param, err, time.Since(start))
	}()
	res, err = local.Services.Accounts.Delete(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Accounts.Delete returned nil, nil")
	}
	return
}

type wrappedAnnotations struct{}

func (s wrappedAnnotations) List(ctx context.Context, param *sourcegraph.AnnotationsListOptions) (res *sourcegraph.AnnotationList, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Annotations", "List", param)
	defer func() {
		trace.After(ctx, "Annotations", "List", param, err, time.Since(start))
	}()
	res, err = local.Services.Annotations.List(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Annotations.List returned nil, nil")
	}
	return
}

type wrappedAuth struct{}

func (s wrappedAuth) GetAccessToken(ctx context.Context, param *sourcegraph.AccessTokenRequest) (res *sourcegraph.AccessTokenResponse, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Auth", "GetAccessToken", param)
	defer func() {
		trace.After(ctx, "Auth", "GetAccessToken", param, err, time.Since(start))
	}()
	res, err = local.Services.Auth.GetAccessToken(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Auth.GetAccessToken returned nil, nil")
	}
	return
}

func (s wrappedAuth) Identify(ctx context.Context, param *pbtypes.Void) (res *sourcegraph.AuthInfo, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Auth", "Identify", param)
	defer func() {
		trace.After(ctx, "Auth", "Identify", param, err, time.Since(start))
	}()
	res, err = local.Services.Auth.Identify(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Auth.Identify returned nil, nil")
	}
	return
}

func (s wrappedAuth) GetExternalToken(ctx context.Context, param *sourcegraph.ExternalTokenSpec) (res *sourcegraph.ExternalToken, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Auth", "GetExternalToken", param)
	defer func() {
		trace.After(ctx, "Auth", "GetExternalToken", param, err, time.Since(start))
	}()
	res, err = local.Services.Auth.GetExternalToken(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Auth.GetExternalToken returned nil, nil")
	}
	return
}

func (s wrappedAuth) SetExternalToken(ctx context.Context, param *sourcegraph.ExternalToken) (res *pbtypes.Void, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Auth", "SetExternalToken", param)
	defer func() {
		trace.After(ctx, "Auth", "SetExternalToken", param, err, time.Since(start))
	}()
	res, err = local.Services.Auth.SetExternalToken(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Auth.SetExternalToken returned nil, nil")
	}
	return
}

type wrappedBuilds struct{}

func (s wrappedBuilds) Get(ctx context.Context, param *sourcegraph.BuildSpec) (res *sourcegraph.Build, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Builds", "Get", param)
	defer func() {
		trace.After(ctx, "Builds", "Get", param, err, time.Since(start))
	}()
	res, err = local.Services.Builds.Get(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Builds.Get returned nil, nil")
	}
	return
}

func (s wrappedBuilds) List(ctx context.Context, param *sourcegraph.BuildListOptions) (res *sourcegraph.BuildList, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Builds", "List", param)
	defer func() {
		trace.After(ctx, "Builds", "List", param, err, time.Since(start))
	}()
	res, err = local.Services.Builds.List(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Builds.List returned nil, nil")
	}
	return
}

func (s wrappedBuilds) Create(ctx context.Context, param *sourcegraph.BuildsCreateOp) (res *sourcegraph.Build, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Builds", "Create", param)
	defer func() {
		trace.After(ctx, "Builds", "Create", param, err, time.Since(start))
	}()
	res, err = local.Services.Builds.Create(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Builds.Create returned nil, nil")
	}
	return
}

func (s wrappedBuilds) Update(ctx context.Context, param *sourcegraph.BuildsUpdateOp) (res *sourcegraph.Build, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Builds", "Update", param)
	defer func() {
		trace.After(ctx, "Builds", "Update", param, err, time.Since(start))
	}()
	res, err = local.Services.Builds.Update(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Builds.Update returned nil, nil")
	}
	return
}

func (s wrappedBuilds) ListBuildTasks(ctx context.Context, param *sourcegraph.BuildsListBuildTasksOp) (res *sourcegraph.BuildTaskList, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Builds", "ListBuildTasks", param)
	defer func() {
		trace.After(ctx, "Builds", "ListBuildTasks", param, err, time.Since(start))
	}()
	res, err = local.Services.Builds.ListBuildTasks(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Builds.ListBuildTasks returned nil, nil")
	}
	return
}

func (s wrappedBuilds) CreateTasks(ctx context.Context, param *sourcegraph.BuildsCreateTasksOp) (res *sourcegraph.BuildTaskList, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Builds", "CreateTasks", param)
	defer func() {
		trace.After(ctx, "Builds", "CreateTasks", param, err, time.Since(start))
	}()
	res, err = local.Services.Builds.CreateTasks(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Builds.CreateTasks returned nil, nil")
	}
	return
}

func (s wrappedBuilds) UpdateTask(ctx context.Context, param *sourcegraph.BuildsUpdateTaskOp) (res *sourcegraph.BuildTask, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Builds", "UpdateTask", param)
	defer func() {
		trace.After(ctx, "Builds", "UpdateTask", param, err, time.Since(start))
	}()
	res, err = local.Services.Builds.UpdateTask(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Builds.UpdateTask returned nil, nil")
	}
	return
}

func (s wrappedBuilds) GetTaskLog(ctx context.Context, param *sourcegraph.BuildsGetTaskLogOp) (res *sourcegraph.LogEntries, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Builds", "GetTaskLog", param)
	defer func() {
		trace.After(ctx, "Builds", "GetTaskLog", param, err, time.Since(start))
	}()
	res, err = local.Services.Builds.GetTaskLog(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Builds.GetTaskLog returned nil, nil")
	}
	return
}

func (s wrappedBuilds) DequeueNext(ctx context.Context, param *sourcegraph.BuildsDequeueNextOp) (res *sourcegraph.Build, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Builds", "DequeueNext", param)
	defer func() {
		trace.After(ctx, "Builds", "DequeueNext", param, err, time.Since(start))
	}()
	res, err = local.Services.Builds.DequeueNext(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Builds.DequeueNext returned nil, nil")
	}
	return
}

type wrappedDefs struct{}

func (s wrappedDefs) Get(ctx context.Context, param *sourcegraph.DefsGetOp) (res *sourcegraph.Def, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Defs", "Get", param)
	defer func() {
		trace.After(ctx, "Defs", "Get", param, err, time.Since(start))
	}()
	res, err = local.Services.Defs.Get(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Defs.Get returned nil, nil")
	}
	return
}

func (s wrappedDefs) List(ctx context.Context, param *sourcegraph.DefListOptions) (res *sourcegraph.DefList, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Defs", "List", param)
	defer func() {
		trace.After(ctx, "Defs", "List", param, err, time.Since(start))
	}()
	res, err = local.Services.Defs.List(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Defs.List returned nil, nil")
	}
	return
}

func (s wrappedDefs) ListRefs(ctx context.Context, param *sourcegraph.DefsListRefsOp) (res *sourcegraph.RefList, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Defs", "ListRefs", param)
	defer func() {
		trace.After(ctx, "Defs", "ListRefs", param, err, time.Since(start))
	}()
	res, err = local.Services.Defs.ListRefs(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Defs.ListRefs returned nil, nil")
	}
	return
}

func (s wrappedDefs) ListRefLocations(ctx context.Context, param *sourcegraph.DefsListRefLocationsOp) (res *sourcegraph.RefLocationsList, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Defs", "ListRefLocations", param)
	defer func() {
		trace.After(ctx, "Defs", "ListRefLocations", param, err, time.Since(start))
	}()
	res, err = local.Services.Defs.ListRefLocations(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Defs.ListRefLocations returned nil, nil")
	}
	return
}

func (s wrappedDefs) ListAuthors(ctx context.Context, param *sourcegraph.DefsListAuthorsOp) (res *sourcegraph.DefAuthorList, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Defs", "ListAuthors", param)
	defer func() {
		trace.After(ctx, "Defs", "ListAuthors", param, err, time.Since(start))
	}()
	res, err = local.Services.Defs.ListAuthors(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Defs.ListAuthors returned nil, nil")
	}
	return
}

type wrappedDeltas struct{}

func (s wrappedDeltas) Get(ctx context.Context, param *sourcegraph.DeltaSpec) (res *sourcegraph.Delta, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Deltas", "Get", param)
	defer func() {
		trace.After(ctx, "Deltas", "Get", param, err, time.Since(start))
	}()
	res, err = local.Services.Deltas.Get(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Deltas.Get returned nil, nil")
	}
	return
}

func (s wrappedDeltas) ListFiles(ctx context.Context, param *sourcegraph.DeltasListFilesOp) (res *sourcegraph.DeltaFiles, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Deltas", "ListFiles", param)
	defer func() {
		trace.After(ctx, "Deltas", "ListFiles", param, err, time.Since(start))
	}()
	res, err = local.Services.Deltas.ListFiles(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Deltas.ListFiles returned nil, nil")
	}
	return
}

type wrappedMeta struct{}

func (s wrappedMeta) Status(ctx context.Context, param *pbtypes.Void) (res *sourcegraph.ServerStatus, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Meta", "Status", param)
	defer func() {
		trace.After(ctx, "Meta", "Status", param, err, time.Since(start))
	}()
	res, err = local.Services.Meta.Status(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Meta.Status returned nil, nil")
	}
	return
}

func (s wrappedMeta) Config(ctx context.Context, param *pbtypes.Void) (res *sourcegraph.ServerConfig, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Meta", "Config", param)
	defer func() {
		trace.After(ctx, "Meta", "Config", param, err, time.Since(start))
	}()
	res, err = local.Services.Meta.Config(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Meta.Config returned nil, nil")
	}
	return
}

type wrappedMirrorRepos struct{}

func (s wrappedMirrorRepos) RefreshVCS(ctx context.Context, param *sourcegraph.MirrorReposRefreshVCSOp) (res *pbtypes.Void, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "MirrorRepos", "RefreshVCS", param)
	defer func() {
		trace.After(ctx, "MirrorRepos", "RefreshVCS", param, err, time.Since(start))
	}()
	res, err = local.Services.MirrorRepos.RefreshVCS(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "MirrorRepos.RefreshVCS returned nil, nil")
	}
	return
}

type wrappedNotify struct{}

func (s wrappedNotify) GenericEvent(ctx context.Context, param *sourcegraph.NotifyGenericEvent) (res *pbtypes.Void, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Notify", "GenericEvent", param)
	defer func() {
		trace.After(ctx, "Notify", "GenericEvent", param, err, time.Since(start))
	}()
	res, err = local.Services.Notify.GenericEvent(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Notify.GenericEvent returned nil, nil")
	}
	return
}

type wrappedOrgs struct{}

func (s wrappedOrgs) Get(ctx context.Context, param *sourcegraph.OrgSpec) (res *sourcegraph.Org, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Orgs", "Get", param)
	defer func() {
		trace.After(ctx, "Orgs", "Get", param, err, time.Since(start))
	}()
	res, err = local.Services.Orgs.Get(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Orgs.Get returned nil, nil")
	}
	return
}

func (s wrappedOrgs) List(ctx context.Context, param *sourcegraph.OrgsListOp) (res *sourcegraph.OrgList, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Orgs", "List", param)
	defer func() {
		trace.After(ctx, "Orgs", "List", param, err, time.Since(start))
	}()
	res, err = local.Services.Orgs.List(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Orgs.List returned nil, nil")
	}
	return
}

func (s wrappedOrgs) ListMembers(ctx context.Context, param *sourcegraph.OrgsListMembersOp) (res *sourcegraph.UserList, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Orgs", "ListMembers", param)
	defer func() {
		trace.After(ctx, "Orgs", "ListMembers", param, err, time.Since(start))
	}()
	res, err = local.Services.Orgs.ListMembers(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Orgs.ListMembers returned nil, nil")
	}
	return
}

type wrappedPeople struct{}

func (s wrappedPeople) Get(ctx context.Context, param *sourcegraph.PersonSpec) (res *sourcegraph.Person, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "People", "Get", param)
	defer func() {
		trace.After(ctx, "People", "Get", param, err, time.Since(start))
	}()
	res, err = local.Services.People.Get(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "People.Get returned nil, nil")
	}
	return
}

type wrappedRepoStatuses struct{}

func (s wrappedRepoStatuses) GetCombined(ctx context.Context, param *sourcegraph.RepoRevSpec) (res *sourcegraph.CombinedStatus, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "RepoStatuses", "GetCombined", param)
	defer func() {
		trace.After(ctx, "RepoStatuses", "GetCombined", param, err, time.Since(start))
	}()
	res, err = local.Services.RepoStatuses.GetCombined(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "RepoStatuses.GetCombined returned nil, nil")
	}
	return
}

func (s wrappedRepoStatuses) Create(ctx context.Context, param *sourcegraph.RepoStatusesCreateOp) (res *sourcegraph.RepoStatus, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "RepoStatuses", "Create", param)
	defer func() {
		trace.After(ctx, "RepoStatuses", "Create", param, err, time.Since(start))
	}()
	res, err = local.Services.RepoStatuses.Create(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "RepoStatuses.Create returned nil, nil")
	}
	return
}

type wrappedRepoTree struct{}

func (s wrappedRepoTree) Get(ctx context.Context, param *sourcegraph.RepoTreeGetOp) (res *sourcegraph.TreeEntry, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "RepoTree", "Get", param)
	defer func() {
		trace.After(ctx, "RepoTree", "Get", param, err, time.Since(start))
	}()
	res, err = local.Services.RepoTree.Get(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "RepoTree.Get returned nil, nil")
	}
	return
}

func (s wrappedRepoTree) Search(ctx context.Context, param *sourcegraph.RepoTreeSearchOp) (res *sourcegraph.VCSSearchResultList, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "RepoTree", "Search", param)
	defer func() {
		trace.After(ctx, "RepoTree", "Search", param, err, time.Since(start))
	}()
	res, err = local.Services.RepoTree.Search(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "RepoTree.Search returned nil, nil")
	}
	return
}

func (s wrappedRepoTree) List(ctx context.Context, param *sourcegraph.RepoTreeListOp) (res *sourcegraph.RepoTreeListResult, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "RepoTree", "List", param)
	defer func() {
		trace.After(ctx, "RepoTree", "List", param, err, time.Since(start))
	}()
	res, err = local.Services.RepoTree.List(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "RepoTree.List returned nil, nil")
	}
	return
}

type wrappedRepos struct{}

func (s wrappedRepos) Get(ctx context.Context, param *sourcegraph.RepoSpec) (res *sourcegraph.Repo, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Repos", "Get", param)
	defer func() {
		trace.After(ctx, "Repos", "Get", param, err, time.Since(start))
	}()
	res, err = local.Services.Repos.Get(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Repos.Get returned nil, nil")
	}
	return
}

func (s wrappedRepos) Resolve(ctx context.Context, param *sourcegraph.RepoResolveOp) (res *sourcegraph.RepoResolution, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Repos", "Resolve", param)
	defer func() {
		trace.After(ctx, "Repos", "Resolve", param, err, time.Since(start))
	}()
	res, err = local.Services.Repos.Resolve(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Repos.Resolve returned nil, nil")
	}
	return
}

func (s wrappedRepos) List(ctx context.Context, param *sourcegraph.RepoListOptions) (res *sourcegraph.RepoList, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Repos", "List", param)
	defer func() {
		trace.After(ctx, "Repos", "List", param, err, time.Since(start))
	}()
	res, err = local.Services.Repos.List(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Repos.List returned nil, nil")
	}
	return
}

func (s wrappedRepos) ListRemote(ctx context.Context, param *sourcegraph.ReposListRemoteOptions) (res *sourcegraph.RemoteRepoList, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Repos", "ListRemote", param)
	defer func() {
		trace.After(ctx, "Repos", "ListRemote", param, err, time.Since(start))
	}()
	res, err = local.Services.Repos.ListRemote(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Repos.ListRemote returned nil, nil")
	}
	return
}

func (s wrappedRepos) Create(ctx context.Context, param *sourcegraph.ReposCreateOp) (res *sourcegraph.Repo, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Repos", "Create", param)
	defer func() {
		trace.After(ctx, "Repos", "Create", param, err, time.Since(start))
	}()
	res, err = local.Services.Repos.Create(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Repos.Create returned nil, nil")
	}
	return
}

func (s wrappedRepos) Update(ctx context.Context, param *sourcegraph.ReposUpdateOp) (res *sourcegraph.Repo, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Repos", "Update", param)
	defer func() {
		trace.After(ctx, "Repos", "Update", param, err, time.Since(start))
	}()
	res, err = local.Services.Repos.Update(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Repos.Update returned nil, nil")
	}
	return
}

func (s wrappedRepos) Delete(ctx context.Context, param *sourcegraph.RepoSpec) (res *pbtypes.Void, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Repos", "Delete", param)
	defer func() {
		trace.After(ctx, "Repos", "Delete", param, err, time.Since(start))
	}()
	res, err = local.Services.Repos.Delete(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Repos.Delete returned nil, nil")
	}
	return
}

func (s wrappedRepos) GetConfig(ctx context.Context, param *sourcegraph.RepoSpec) (res *sourcegraph.RepoConfig, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Repos", "GetConfig", param)
	defer func() {
		trace.After(ctx, "Repos", "GetConfig", param, err, time.Since(start))
	}()
	res, err = local.Services.Repos.GetConfig(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Repos.GetConfig returned nil, nil")
	}
	return
}

func (s wrappedRepos) GetCommit(ctx context.Context, param *sourcegraph.RepoRevSpec) (res *vcs.Commit, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Repos", "GetCommit", param)
	defer func() {
		trace.After(ctx, "Repos", "GetCommit", param, err, time.Since(start))
	}()
	res, err = local.Services.Repos.GetCommit(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Repos.GetCommit returned nil, nil")
	}
	return
}

func (s wrappedRepos) ListCommits(ctx context.Context, param *sourcegraph.ReposListCommitsOp) (res *sourcegraph.CommitList, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Repos", "ListCommits", param)
	defer func() {
		trace.After(ctx, "Repos", "ListCommits", param, err, time.Since(start))
	}()
	res, err = local.Services.Repos.ListCommits(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Repos.ListCommits returned nil, nil")
	}
	return
}

func (s wrappedRepos) ListBranches(ctx context.Context, param *sourcegraph.ReposListBranchesOp) (res *sourcegraph.BranchList, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Repos", "ListBranches", param)
	defer func() {
		trace.After(ctx, "Repos", "ListBranches", param, err, time.Since(start))
	}()
	res, err = local.Services.Repos.ListBranches(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Repos.ListBranches returned nil, nil")
	}
	return
}

func (s wrappedRepos) ListTags(ctx context.Context, param *sourcegraph.ReposListTagsOp) (res *sourcegraph.TagList, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Repos", "ListTags", param)
	defer func() {
		trace.After(ctx, "Repos", "ListTags", param, err, time.Since(start))
	}()
	res, err = local.Services.Repos.ListTags(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Repos.ListTags returned nil, nil")
	}
	return
}

func (s wrappedRepos) ListCommitters(ctx context.Context, param *sourcegraph.ReposListCommittersOp) (res *sourcegraph.CommitterList, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Repos", "ListCommitters", param)
	defer func() {
		trace.After(ctx, "Repos", "ListCommitters", param, err, time.Since(start))
	}()
	res, err = local.Services.Repos.ListCommitters(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Repos.ListCommitters returned nil, nil")
	}
	return
}

func (s wrappedRepos) GetSrclibDataVersionForPath(ctx context.Context, param *sourcegraph.TreeEntrySpec) (res *sourcegraph.SrclibDataVersion, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Repos", "GetSrclibDataVersionForPath", param)
	defer func() {
		trace.After(ctx, "Repos", "GetSrclibDataVersionForPath", param, err, time.Since(start))
	}()
	res, err = local.Services.Repos.GetSrclibDataVersionForPath(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Repos.GetSrclibDataVersionForPath returned nil, nil")
	}
	return
}

func (s wrappedRepos) ConfigureApp(ctx context.Context, param *sourcegraph.RepoConfigureAppOp) (res *pbtypes.Void, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Repos", "ConfigureApp", param)
	defer func() {
		trace.After(ctx, "Repos", "ConfigureApp", param, err, time.Since(start))
	}()
	res, err = local.Services.Repos.ConfigureApp(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Repos.ConfigureApp returned nil, nil")
	}
	return
}

func (s wrappedRepos) GetInventory(ctx context.Context, param *sourcegraph.RepoRevSpec) (res *inventory.Inventory, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Repos", "GetInventory", param)
	defer func() {
		trace.After(ctx, "Repos", "GetInventory", param, err, time.Since(start))
	}()
	res, err = local.Services.Repos.GetInventory(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Repos.GetInventory returned nil, nil")
	}
	return
}

func (s wrappedRepos) ReceivePack(ctx context.Context, param *sourcegraph.ReceivePackOp) (res *sourcegraph.Packet, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Repos", "ReceivePack", param)
	defer func() {
		trace.After(ctx, "Repos", "ReceivePack", param, err, time.Since(start))
	}()
	res, err = local.Services.Repos.ReceivePack(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Repos.ReceivePack returned nil, nil")
	}
	return
}

func (s wrappedRepos) UploadPack(ctx context.Context, param *sourcegraph.UploadPackOp) (res *sourcegraph.Packet, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Repos", "UploadPack", param)
	defer func() {
		trace.After(ctx, "Repos", "UploadPack", param, err, time.Since(start))
	}()
	res, err = local.Services.Repos.UploadPack(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Repos.UploadPack returned nil, nil")
	}
	return
}

type wrappedSearch struct{}

func (s wrappedSearch) Search(ctx context.Context, param *sourcegraph.SearchOp) (res *sourcegraph.SearchResultsList, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Search", "Search", param)
	defer func() {
		trace.After(ctx, "Search", "Search", param, err, time.Since(start))
	}()
	res, err = local.Services.Search.Search(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Search.Search returned nil, nil")
	}
	return
}

func (s wrappedSearch) RefreshIndex(ctx context.Context, param *sourcegraph.SearchRefreshIndexOp) (res *pbtypes.Void, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Search", "RefreshIndex", param)
	defer func() {
		trace.After(ctx, "Search", "RefreshIndex", param, err, time.Since(start))
	}()
	res, err = local.Services.Search.RefreshIndex(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Search.RefreshIndex returned nil, nil")
	}
	return
}

type wrappedUsers struct{}

func (s wrappedUsers) Get(ctx context.Context, param *sourcegraph.UserSpec) (res *sourcegraph.User, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Users", "Get", param)
	defer func() {
		trace.After(ctx, "Users", "Get", param, err, time.Since(start))
	}()
	res, err = local.Services.Users.Get(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Users.Get returned nil, nil")
	}
	return
}

func (s wrappedUsers) GetWithEmail(ctx context.Context, param *sourcegraph.EmailAddr) (res *sourcegraph.User, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Users", "GetWithEmail", param)
	defer func() {
		trace.After(ctx, "Users", "GetWithEmail", param, err, time.Since(start))
	}()
	res, err = local.Services.Users.GetWithEmail(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Users.GetWithEmail returned nil, nil")
	}
	return
}

func (s wrappedUsers) ListEmails(ctx context.Context, param *sourcegraph.UserSpec) (res *sourcegraph.EmailAddrList, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Users", "ListEmails", param)
	defer func() {
		trace.After(ctx, "Users", "ListEmails", param, err, time.Since(start))
	}()
	res, err = local.Services.Users.ListEmails(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Users.ListEmails returned nil, nil")
	}
	return
}

func (s wrappedUsers) List(ctx context.Context, param *sourcegraph.UsersListOptions) (res *sourcegraph.UserList, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Users", "List", param)
	defer func() {
		trace.After(ctx, "Users", "List", param, err, time.Since(start))
	}()
	res, err = local.Services.Users.List(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Users.List returned nil, nil")
	}
	return
}

func (s wrappedUsers) ListTeammates(ctx context.Context, param *sourcegraph.UserSpec) (res *sourcegraph.Teammates, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Users", "ListTeammates", param)
	defer func() {
		trace.After(ctx, "Users", "ListTeammates", param, err, time.Since(start))
	}()
	res, err = local.Services.Users.ListTeammates(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Users.ListTeammates returned nil, nil")
	}
	return
}

func (s wrappedUsers) Count(ctx context.Context, param *pbtypes.Void) (res *sourcegraph.UserCount, err error) {
	start := time.Now()
	ctx = trace.Before(ctx, "Users", "Count", param)
	defer func() {
		trace.After(ctx, "Users", "Count", param, err, time.Since(start))
	}()
	res, err = local.Services.Users.Count(ctx, param)
	if res == nil && err == nil {
		err = grpc.Errorf(codes.Internal, "Users.Count returned nil, nil")
	}
	return
}
