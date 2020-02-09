/**Here the component renders two buttons, for logging in and logging out, 
 * depending on whether the user is currently authenticated. */

import React from "react";
import { useAuth0 } from "../react-auth0-spa"
import { Link } from "react-router-dom"
import Navbar from 'react-bootstrap/Navbar'
import Form from 'react-bootstrap/Form'
import FormControl from 'react-bootstrap/FormControl'
import Button from 'react-bootstrap/Button'

const NavBar = () => {
  const { isAuthenticated, loginWithRedirect, logout } = useAuth0();

  return (
    <Navbar bg="light" expand="lg" className="bg-light justify-content-between">
      <Navbar.Brand href="#home">Quedat al barri</Navbar.Brand>
    <div>
    {isAuthenticated && (
      <span>
        <Link to="/">Home</Link>&nbsp;
        <Link to="/profile">Profile</Link>
      </span>
    )}
    </div>
    <Form inline>
        {!isAuthenticated && (
          <Button onClick={() => loginWithRedirect({})}>Log in</Button>
        )}

        {isAuthenticated && <Button onClick={() => logout()}>Log out</Button>}
      </Form>
    </Navbar>
  );
};

export default NavBar;