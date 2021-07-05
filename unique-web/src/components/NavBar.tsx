import React from "react";
import "../style/NavBar.scss";

export class NavBar extends React.Component {
  render() {
    return (
      <div className="navBar">
        <ul className="navBarUl">
          <li>
            <a href="/">Home</a>
          </li>
          <li>
            <a href="/user/profile/">My Profile</a>
          </li>
          <li id="navLogout">
            <a href="/logout">Logout</a>
          </li>
        </ul>
      </div>
    );
  }
}
