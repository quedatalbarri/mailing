// src/App.js

import React from "react";
import NavBar from "./components/NavBar";

// New - import the React Router components, and the Profile page component
import { Router, Route, Switch } from "react-router-dom";
import Profile from "./components/Profile";
import Barris from "./components/Barris";
import history from "./utils/history";
import PrivateRoute from "./components/PrivateRoute"
import Container from 'react-bootstrap/Container'

function App() {
  return (
    <div className="App">
      {/* Don't forget to include the history module */}
      <Router history={history}>
        <header>
          <NavBar />
        </header>
        <Container>
          <Switch>
            {/* <Route path="/" exact /> */}
            <Route path="/" exact component={Barris} />
            <PrivateRoute path="/profile" component={Profile} />
          </Switch>
        </Container>
      </Router>
    </div>
    
  );
}

export default App;