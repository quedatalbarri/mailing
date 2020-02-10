/**This component takes another component as one of its arguments. It makes use of the 
 * useEffect hook to redirect the user to the login page if they are not yet authenticated.
 * If the user is authenticated, the redirect will not take place and the component that was specified 
 * as the argument will be rendered instead. In this way, components that require the user to be logged 
 * in can be protected simply by wrapping the component using PrivateRoute. */

import React, { useEffect } from "react";
import { Route } from "react-router-dom";
import { useAuth0 } from "../react-auth0-spa";

const PrivateRoute = ({ component: Component, path, ...rest }) => {
  const { loading, isAuthenticated, loginWithRedirect } = useAuth0();

  useEffect(() => {
    if (loading || isAuthenticated) {
      return;
    }
    const fn = async () => {
      await loginWithRedirect({
        appState: { targetUrl: path }
      });
    };
    fn();
  }, [loading, isAuthenticated, loginWithRedirect, path]);

  const render = props =>
    isAuthenticated === true ? <Component {...props} /> : null;

  return <Route path={path} render={render} {...rest} />;
};

export default PrivateRoute;