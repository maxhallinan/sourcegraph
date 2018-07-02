import DirectionalSignIcon from '@sourcegraph/icons/lib/DirectionalSign'
import * as React from 'react'
import { Route, RouteComponentProps, Switch } from 'react-router'
import { Redirect } from 'react-router-dom'
import { Subscription } from 'rxjs'
import { map } from 'rxjs/operators'
import * as GQL from '../../backend/graphqlschema'
import { HeroPage } from '../../components/HeroPage'
import { siteFlags } from '../../site/backend'
import { UserAreaPageProps } from '../area/UserArea'
import { UserSettingsAccountPage } from './UserSettingsAccountPage'
import { UserSettingsConfigurationPage } from './UserSettingsConfigurationPage'
import { UserSettingsCreateAccessTokenPage } from './UserSettingsCreateAccessTokenPage'
import { UserSettingsEmailsPage } from './UserSettingsEmailsPage'
import { UserSettingsExternalAccountsPage } from './UserSettingsExternalAccountsPage'
import { UserSettingsIntegrationsPage } from './UserSettingsIntegrationsPage'
import { UserSettingsProfilePage } from './UserSettingsProfilePage'
import { UserSettingsSidebar } from './UserSettingsSidebar'
import { UserSettingsTokensPage } from './UserSettingsTokensPage'

const NotFoundPage = () => <HeroPage icon={DirectionalSignIcon} title="404: Not Found" />

interface Props extends UserAreaPageProps, RouteComponentProps<{}> {
    isLightTheme: boolean
    onThemeChange: () => void
}

interface State {
    externalAuthEnabled: boolean

    /**
     * Holds the newly created access token (from UserSettingsCreateAccessTokenPage), if any. After
     * it is displayed to the user, this subject's value is set back to undefined.
     */
    newlyCreatedAccessToken?: GQL.ICreateAccessTokenResult
}

/**
 * Renders a layout of a sidebar and a content area to display user settings.
 */
export class UserSettingsArea extends React.Component<Props, State> {
    public state: State = { externalAuthEnabled: false }
    private subscriptions = new Subscription()

    public componentDidMount(): void {
        this.subscriptions.add(
            siteFlags.pipe(map(({ externalAuthEnabled }) => externalAuthEnabled)).subscribe(externalAuthEnabled => {
                this.setState({ externalAuthEnabled })
            })
        )
    }

    public componentWillUnmount(): void {
        this.subscriptions.unsubscribe()
    }

    public render(): JSX.Element | null {
        if (!this.props.user) {
            return null
        }

        // If not logged in, redirect to sign in
        if (!this.props.authenticatedUser) {
            const newUrl = new URL(window.location.href)
            newUrl.pathname = '/sign-in'
            // Return to the current page after sign up/in.
            newUrl.searchParams.set('returnTo', window.location.href)
            return <Redirect to={newUrl.pathname + newUrl.search} />
        }

        if (this.props.authenticatedUser.id !== this.props.user.id && !this.props.user.viewerCanAdminister) {
            return (
                <HeroPage
                    icon={DirectionalSignIcon}
                    title="403: Forbidden"
                    subtitle="You are not authorized to view or edit this user's settings."
                />
            )
        }

        if (this.props.match.isExact) {
            return <Redirect to={`${this.props.match.path}/profile`} />
        }

        return (
            <div className="user-settings-area area">
                <UserSettingsSidebar
                    externalAuthEnabled={this.state.externalAuthEnabled}
                    className="area__sidebar"
                    {...this.props}
                />
                <div className="area__content">
                    <Switch>
                        {/* Render empty page if no settings page selected */}
                        <Route
                            path={`${this.props.match.url}/profile`}
                            key="hardcoded-key" // see https://github.com/ReactTraining/react-router/issues/4578#issuecomment-334489490
                            exact={true}
                            // tslint:disable-next-line:jsx-no-lambda
                            render={routeComponentProps => (
                                <UserSettingsProfilePage {...routeComponentProps} {...this.props} />
                            )}
                        />
                        <Route
                            path={`${this.props.match.url}/configuration`}
                            key="hardcoded-key" // see https://github.com/ReactTraining/react-router/issues/4578#issuecomment-334489490
                            exact={true}
                            // tslint:disable-next-line:jsx-no-lambda
                            render={routeComponentProps => (
                                <UserSettingsConfigurationPage {...routeComponentProps} {...this.props} />
                            )}
                        />
                        {!this.state.externalAuthEnabled && (
                            <Route
                                path={`${this.props.match.url}/account`}
                                key="hardcoded-key" // see https://github.com/ReactTraining/react-router/issues/4578#issuecomment-334489490
                                exact={true}
                                // tslint:disable-next-line:jsx-no-lambda
                                render={routeComponentProps => (
                                    <UserSettingsAccountPage {...routeComponentProps} {...this.props} />
                                )}
                            />
                        )}
                        <Route
                            path={`${this.props.match.url}/emails`}
                            key="hardcoded-key" // see https://github.com/ReactTraining/react-router/issues/4578#issuecomment-334489490
                            exact={true}
                            // tslint:disable-next-line:jsx-no-lambda
                            render={routeComponentProps => (
                                <UserSettingsEmailsPage {...routeComponentProps} {...this.props} />
                            )}
                        />
                        <Route
                            path={`${this.props.match.url}/external-accounts`}
                            key="hardcoded-key" // see https://github.com/ReactTraining/react-router/issues/4578#issuecomment-334489490
                            exact={true}
                            // tslint:disable-next-line:jsx-no-lambda
                            render={routeComponentProps => (
                                <UserSettingsExternalAccountsPage {...routeComponentProps} {...this.props} />
                            )}
                        />
                        {window.context.accessTokensEnabled && (
                            <Route
                                path={`${this.props.match.url}/tokens`}
                                key="hardcoded-key" // see https://github.com/ReactTraining/react-router/issues/4578#issuecomment-334489490
                                exact={true}
                                // tslint:disable-next-line:jsx-no-lambda
                                render={routeComponentProps => (
                                    <UserSettingsTokensPage
                                        {...routeComponentProps}
                                        {...this.props}
                                        newToken={this.state.newlyCreatedAccessToken}
                                        onDidPresentNewToken={this.setNewToken}
                                    />
                                )}
                            />
                        )}
                        {window.context.accessTokensEnabled && (
                            <Route
                                path={`${this.props.match.url}/tokens/new`}
                                key="hardcoded-key" // see https://github.com/ReactTraining/react-router/issues/4578#issuecomment-334489490
                                exact={true}
                                // tslint:disable-next-line:jsx-no-lambda
                                render={routeComponentProps => (
                                    <UserSettingsCreateAccessTokenPage
                                        {...routeComponentProps}
                                        {...this.props}
                                        onDidCreateAccessToken={this.setNewToken}
                                    />
                                )}
                            />
                        )}
                        <Route
                            path={`${this.props.match.url}/integrations`}
                            key="hardcoded-key"
                            component={UserSettingsIntegrationsPage}
                            exact={true}
                        />
                        <Route component={NotFoundPage} key="hardcoded-key" />
                    </Switch>
                </div>
            </div>
        )
    }

    private setNewToken = (value?: GQL.ICreateAccessTokenResult): void => {
        this.setState({ newlyCreatedAccessToken: value })
    }
}