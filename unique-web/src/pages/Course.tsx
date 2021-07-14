import React from "react";
import { Redirect, RouteComponentProps } from "react-router-dom";
import { getCourseFromId } from "../apiCommunication";
import { NavBar } from "../components/NavBar";
import { setTitle } from "../utils";

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

  courseInfo: CourseInfo | null;
};

export class Course extends React.Component<CourseProps, CourseState> {
  state: CourseState = {
    isLoaded: false,
    error: null,
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
      console.log(courseInfo);
      return (
        <div className="coursePage">
          <NavBar />

          <h1>{courseInfo.title}</h1>
          <h3>{courseInfo.subtitle}</h3>
        </div>
      );
    } else {
      return <Redirect to="/" />;
    }
  }
}
