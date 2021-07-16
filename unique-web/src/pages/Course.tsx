import React from "react";
import { Redirect, RouteComponentProps } from "react-router-dom";
import { getCourseFromId } from "../apiCommunication";
import { NavBar } from "../components/NavBar";
import { setTitle } from "../utils";
import "../style/Course.scss";
import { cookies } from "..";

interface CourseProps
  extends RouteComponentProps<{ cid: string | undefined }> {}

type CourseInfo = {
  id: string;
  title: string;
  subtitle: string;
  description: string;

  teachers_id: string[];
  students_id: string[];
};

type CourseState = {
  isLoaded: boolean;
  error: string | null;
  user: string | undefined;
  courseInfo: CourseInfo | null;
};

export class Course extends React.Component<CourseProps, CourseState> {
  state: CourseState = {
    isLoaded: false,
    error: null,
    user: undefined,
    courseInfo: null,
  };

  componentDidMount() {
    const { cid } = this.props.match.params;
    if (cid === undefined) {
      this.setState({
        isLoaded: true,
        error: "No course ID provided.",
      });
    } else {
      getCourseFromId(cid).then(
        (result) => {
          if (result["error"] !== undefined) {
            this.setState({
              isLoaded: true,
              error: result["error"],
            });
          } else {
            this.setState({
              isLoaded: true,
              courseInfo: {
                id: cid,
                title: result["title"],
                subtitle: result["subtitle"],
                description: result["description"],
                teachers_id: result["teachers_id"],
                students_id: result["students_id"],
              },
            });
            setTitle(result["title"]);

            const uid = cookies.get("luid");
            if (uid !== undefined) {
              if (result["teachers_id"].indexOf(uid) >= 0) {
                this.setState({ user: "teacher" });
              } else if (result["students_id"].indexOf(uid) >= 0) {
                this.setState({ user: "student" });
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
  }

  render() {
    const { isLoaded, error, courseInfo } = this.state;

    if (!isLoaded) {
      return <h1>Loading your informations...</h1>;
    } else if (error !== null) {
      return <Redirect to={"/error?err" + error} />;
    } else if (courseInfo !== null) {
      return (
        <div className="coursePage">
          <NavBar />

          <section>
            <aside className="courseSidebar">
              <h1 id="courseTitle">{courseInfo.title}</h1>
              <p id="courseSubtitle">{courseInfo.subtitle}</p>

              <hr className="separator" />

              <ul>
                <li>
                  <p>Home</p>
                </li>
                <li>
                  <p>Members</p>
                </li>
                <li>
                  <p>Materials</p>
                </li>
                <li>
                  <p>Assignments</p>
                </li>
                {this.state.user === "teacher" ? (
                  <div id="sidebarTeacherOpts">
                    <hr className="separator" />
                    <li>
                      <p>Course Settings</p>
                    </li>
                    <li>
                      <p>Manage Members</p>
                    </li>
                  </div>
                ) : null}
              </ul>
            </aside>
          </section>
        </div>
      );
    } else {
      return <Redirect to="/" />;
    }
  }
}
