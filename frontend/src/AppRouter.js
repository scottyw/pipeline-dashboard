import React from "react";
import { BrowserRouter as Router, Route, Redirect } from "react-router-dom";

import Index from './Components/Index';
import Timeline from './Components/Timeline';

import { Sidebar, Content } from '@puppet/react-components';

class AppRouter extends React.Component {

  constructor(props) {
    super(props);

    this.state = {
      redirect: false,
      redirect_to: ""
    }
  }

  render() {
    if (this.state.redirect === true) {
      this.setState({
        redirect: false
      })
      return (<Router>
        <Redirect to={this.state.redirect_to} />
      </Router>);
    }

    return (
      <Router>
      <div style={{ float: 'left', position: 'relative', height: '100vh' }}>
        <Sidebar>
          <Sidebar.Header
            logo="CI Dashboard"
            onClick={() => console.log('logo clicked')}
            ariaLabel="Return to the home page"
          />
          <br />
          <Sidebar.Section>
            <Sidebar.Item onClick={() => { this.setState({ redirect: true, redirect_to: "/" })}} title="Home" icon="home" active />
          </Sidebar.Section>
          <Sidebar.Section label="reports">
            <Sidebar.Item onClick={() => { this.setState({ redirect: true, redirect_to: "/timeline" })}} title="Timeline" icon="code" active />
          </Sidebar.Section>

        </Sidebar>
        </div>
        <div style={{ position: 'relative', height: '100vh' }} className="app-main-content">
        <Content>
          <Route path="/" exact component={Index} />
          <Route path="/timeline" exact component={Timeline} />
        </Content>
        </div>
      </Router>
    );
  }
}

export default AppRouter;

// <Navbar bg="light" expand="lg" className="bg-light justify-content-between">
//   <Navbar.Brand href="/">Home</Navbar.Brand>
//   <Nav className="mr-auto">
//     <Nav.Link href="/">Home</Nav.Link>
//     <Nav.Link href="/timeline">Timeline</Nav.Link>
//   </Nav>
//
//   <Navbar.Collapse className="justify-content-end">
//     <Navbar.Text className="mr-sm-2">
//       Current Time: { new Date().toLocaleString('en-US') }
//     </Navbar.Text>
//   </Navbar.Collapse>
// </Navbar>
// <br />
// <br />
