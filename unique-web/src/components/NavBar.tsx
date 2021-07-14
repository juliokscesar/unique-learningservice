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
          <li>
            <a href="/myCourses">My Courses</a>
          </li>
          <li id="navRighty">
            <a href="/logout">Logout</a>
          </li>
          <li id="navRighty">
            <a href="/accountSettings">Settings</a>
          </li>
        </ul>
      </div>
    );
  }
}
