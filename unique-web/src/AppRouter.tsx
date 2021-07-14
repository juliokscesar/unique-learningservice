import React from "react";
import { BrowserRouter, Switch, Route, Redirect } from "react-router-dom";
import { Toast } from "./components/Toast";
import { Error } from "./components/Error";
import { UserHome } from "./pages/UserHome";
import { Login } from "./pages/Login";
import { Register } from "./pages/Register";
import { cookies } from "./index";
import { Profile } from "./pages/Profile";
import { Logout } from "./pages/Logout";
import { UserCourses } from "./pages/UserCourses";
import { CreateCourse } from "./pages/CreateCourse";
import { AccountSettings } from "./pages/AccountSettings";
import { Course } from "./pages/Course";
import { ChangePasswordForm } from "./pages/ChangePasswordForm";

class AppRouter extends React.Component {
  render() {
    const luid = cookies.get("luid");
    const puid = cookies.get("puid");

    const isUserLogged = puid && luid;
    return (
      <div className="App">
        <BrowserRouter>
          <Switch>
            <Route exact path="/">
              <UserHome />
            </Route>

            <Route path="/error" component={Error} />

            <Route path="/login">
              {isUserLogged ? <Redirect to="/" /> : <Login />}
            </Route>

            <Route path="/register">
              {isUserLogged ? <Redirect to="/" /> : <Register />}
            </Route>

            <Route path="/logout">
              <Logout />
            </Route>

            <Route path="/user/profile/:uid" component={Profile} />
            <Route path="/user/profile">
              <Redirect
                to={isUserLogged ? "/user/profile/" + puid : "/login"}
              />
            </Route>

            <Route exact path="/accountSettings">
              {isUserLogged ? <AccountSettings /> : <Redirect to="/login" />}
            </Route>
            <Route path="/accountSettings/changePass">
              {isUserLogged ? <ChangePasswordForm /> : <Redirect to="/login" />}
            </Route>

            <Route path="/myCourses">
              <UserCourses />
            </Route>

            <Route path="/course/:cid" component={Course} />

            <Route path="/createCourse">
              <CreateCourse />
            </Route>

            <Route path="/secret">
              <Toast
                message="Eureka! You have just discovered my secret!"
                author="Julio"
              />
            </Route>

            <Redirect from="*" to="/error?err=Not+Found" />
          </Switch>
        </BrowserRouter>
      </div>
    );
  }
}

export default AppRouter;
