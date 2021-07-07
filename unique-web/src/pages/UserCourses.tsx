import React from "react";
import { Redirect } from "react-router-dom";
import { getManyCoursesFromIds, getUserFromId } from "../apiCommunication";
import { CourseCard } from "../components/CourseCard";
import { NavBar } from "../components/NavBar";
import { cookies } from "../index";
import { setTitle } from "../utils";

type CourseType = {
  id: string;
  title: string;
  subtitle: string;
  description: string;
};

type UserCoursesState = {
  isLoaded: boolean;
  error: string | null;
  loggedUser:
    | {
        name: string;
        courses_id: string[] | null;
      }
    | undefined;
  userCourses: CourseType[];
};

export class UserCourses extends React.Component<{}, UserCoursesState> {
  state: UserCoursesState = {
    isLoaded: false,
    error: null,
    loggedUser: undefined,
    userCourses: [],
  };

  componentDidMount() {
    const userId = cookies.get("luid");
    if (userId === undefined) {
      document.location.href = "/login";
    } else {
      getUserFromId(userId).then(
        (result) => {
          if (result["error"] !== undefined) {
            this.setState({
              isLoaded: true,
              error: result["error"],
            });
          } else {
            this.setState({
              isLoaded: true,
              loggedUser: {
                name: result["name"],
                courses_id: result["courses_id"],
              },
            });

            if (this.state.loggedUser !== undefined) {
              if (this.state.loggedUser.courses_id !== null) {
                getManyCoursesFromIds(this.state.loggedUser.courses_id).then(
                  (result) => {
                    if (result["error"] !== undefined) {
                      this.setState({
                        error: result["error"],
                      });
                    } else {
                      const courses = result.map((c: any) => {
                        return {
                          id: c["id"],
                          title: c["title"],
                          subtitle: c["subtitle"],
                          description: c["description"],
                        };
                      });
                      this.setState({
                        userCourses: courses,
                      });
                    }
                  }
                );
              }
            }
          }
        },
        (err) => {
          this.setState({
            isLoaded: true,
            error: err,
          });
        }
      );
    }

    setTitle("My Courses");
  }

  render() {
    const { isLoaded, error, loggedUser, userCourses } = this.state;

    if (!isLoaded) {
      return <h1>Loading your informations...</h1>;
    } else if (error !== null) {
      return <Redirect to={"/error?err=" + error} />;
    } else if (loggedUser !== undefined) {
      const coursesInfo = userCourses.map((c) => {
        return (
          <CourseCard
            key={c.id}
            title={c.title}
            subtitle={c.subtitle}
            description={c.description}
            courseUrl={"/course/" + c.id}
          />
        );
      });

      return (
        <div className="userCourses">
          <NavBar />

          <h1>{loggedUser.name}'s Courses</h1>

          {coursesInfo}

          <CourseCard
            title="Create"
            subtitle="+"
            description="Create a Course"
            courseUrl={"/createCourse"}
          />
        </div>
      );
    } else {
      return <Redirect to="/login" />;
    }
  }
}
