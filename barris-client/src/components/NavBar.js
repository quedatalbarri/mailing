/**Here the component renders two buttons, for logging in and logging out, 
 * depending on whether the user is currently authenticated. */

import React from "react";
import { useAuth0 } from "../react-auth0-spa"
import { Link } from "react-router-dom"
import Logo from "../images/logo.png"
import Navbar from 'react-bootstrap/Navbar'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import './styles/NavBar.css'

const NavBar = () => {
  const { isAuthenticated, loginWithRedirect, logout } = useAuth0();

  return (
    <Navbar bg="dark" variant="dark" expand="lg" className="justify-content-between">
      <span>

        <Navbar.Brand href="#home" className="qb-navbar-brand">
          <img
            src={Logo}
            width="35"
            height="35"
            className="d-inline-block align-top qb-logo"
            alt="Quedat al barri"
          />
          Quedat al barri</Navbar.Brand>
      </span>
    <div>
    {isAuthenticated && (
      <span>
        <Link to="/" className="qb-link">Home</Link>&nbsp;
        <Link to="/profile" className="qb-link">Profile</Link>
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
  )
}

export default NavBar;