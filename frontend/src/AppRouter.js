import React from "react";
import { BrowserRouter as Router, Route } from "react-router-dom";

import Index from './Components/Index';

import { Navbar } from 'react-bootstrap';
import { Sidebar } from '@puppet/react-components';

function AppRouter() {
  return (
    <Router>
      <Navbar bg="light" expand="lg" className="bg-light justify-content-between">
        <Navbar.Brand href="/">Home</Navbar.Brand>
        <Navbar.Collapse className="justify-content-end">
          <Navbar.Text className="mr-sm-2">
            Current Time: { new Date().toLocaleString('en-US') }
          </Navbar.Text>
        </Navbar.Collapse>
      </Navbar>
      <br />
      <br />
      <div>
        <Sidebar />
        <Route path="/" exact component={Index} />
      </div>
    </Router>
  );
}

export default AppRouter;
