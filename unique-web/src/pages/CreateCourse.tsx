import React from "react";
import { NavBar } from "../components/NavBar";
import { UserForm } from "../components/UserForm";
import { cookies } from "../index";
import { createCourse } from "../apiCommunication";
import { capitalizeFirstLetter } from "../utils";

type CreateCourseState = {
  title: string;
  subtitle: string;
  description: string;
  admId: string;
};

export class CreateCourse extends React.Component<{}, CreateCourseState> {
  state: CreateCourseState = {
    title: "",
    subtitle: "",
    description: "",
    admId: "",
  };

  changeTitle = (newTitle: string) => {
    this.setState({ title: newTitle });
  };

  changeSubtitle = (newSub: string) => {
    this.setState({ subtitle: newSub });
  };

  changeDescription = (newDesc: string) => {
    this.setState({ description: newDesc });
  };

  changeAdmId = (newAdmId: string) => {
    this.setState({ admId: newAdmId });
  };

  createCourseSubmit = async () => {
    return await createCourse(this.state);
  };

  componentDidMount() {
    const userId = cookies.get("luid");
    if (userId === undefined) {
      document.location.href = "/login";
    }

    this.changeAdmId(userId);
  }

  render() {
    const inputs = [
      {
        name: "title",
        type: "text",
        required: true,
        minLength: 3,
        onChangeFn: (e: React.ChangeEvent<HTMLInputElement>) =>
          this.changeTitle(e.target.value),
      },
      {
        name: "subtitle",
        type: "text",
        required: true,
        minLength: 3,
        onChangeFn: (e: React.ChangeEvent<HTMLInputElement>) =>
          this.changeSubtitle(e.target.value),
      },
      {
        name: "description",
        type: "text",
        required: true,
        minLength: 3,
        onChangeFn: (e: React.ChangeEvent<HTMLInputElement>) =>
          this.changeDescription(e.target.value),
      },
    ];

    return (
      <div className="createCourseForm">
        <NavBar />

        <UserForm
          title="Create Course"
          inputs={inputs}
          submitFn={async () => {
            const result = await this.createCourseSubmit();

            if (result["error"] !== undefined) {
              document.location.href =
                "/error?err=" + capitalizeFirstLetter(result["error"]);
            } else {
              document.location.href = "/myCourses";
            }
          }}
          submitDisabled={false}
        />
      </div>
    );
  }
}
