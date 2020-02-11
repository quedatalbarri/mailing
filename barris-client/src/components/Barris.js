/**Notice here that the useAuth0 hook is again being used, this time to retrieve the user's profile information 
 * (through the user property) and a loading property that can be used to display some kind of "loading" text 
 * or spinner graphic while the user's data is being retrieved.
 * In the UI for this component, the user's profile picture, name, and email address is being retrieved
 *  from the user property and displayed on the screen. */

import React, { Fragment } from "react"
import { useAuth0 } from "../react-auth0-spa"
import NewBarri from './NewBarri'

//const config = require('../config.json')
const endpoint = "http://localhost:1323"

const Barris = () => {
  const { loading, user } = useAuth0()

  if (!user) {
    return <div>Registrat</div>;
  }
  else if (loading) {
    return <div>Loading...</div>;
  }
  return (
    <Fragment>
      <h3>Barrios creados</h3>
      <p>Todo- list aqui</p>
      <NewBarri/>
    </Fragment>
  );
};

export default Barris