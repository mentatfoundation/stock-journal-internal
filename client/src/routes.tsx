import React from "react";
import { Link, Route, Switch } from "react-router-dom";
import About from "./components/About";
import Home from "./components/Home";

const Routes = () => {
  return (
    <>
      <Link to="/">Home</Link>
      <Link to="/about">About</Link>
      <Switch>
        <Route exact path="/" component={Home} />
        <Route path="/about" component={About} />
      </Switch>
    </>
  );
};

export default Routes;
