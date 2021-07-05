import React from "react";
import { capitalizeFirstLetter, setTitle } from "../utils";
import "../style/UserForm.scss";

type FnEventChange = (e: React.ChangeEvent<HTMLInputElement>) => void;
type FnSubmit = () => void;

type UserFormProps = {
  title: string;
  inputs: {
    name: string;
    type: string;
    required: boolean;
    minLength?: number;
    onChangeFn?: FnEventChange;
  }[];
  submitFn: FnSubmit;
};

export class UserForm extends React.Component<UserFormProps> {
  render() {
    const inputs = this.props.inputs.map((input) => {
      return (
        //<div className={"input" + capitalizeFirstLetter(input.name)}>
        <input
          type={input.type}
          id={input.name}
          key={input.name}
          name={input.name}
          placeholder={capitalizeFirstLetter(input.name)}
          required={input.required}
          minLength={input.minLength}
          onChange={input.onChangeFn}
        />
        //</div>
      );
    });

    return (
      <div className="userForm">
        <h1>{capitalizeFirstLetter(this.props.title)}</h1>

        {inputs}

        <button id="submit" disabled={true}>
          {capitalizeFirstLetter(this.props.title)}
        </button>
      </div>
    );
  }

  componentDidMount() {
    setTitle(capitalizeFirstLetter(this.props.title));

    const submitBtn = document.querySelector("#submit") as HTMLInputElement;

    for (let input of this.props.inputs) {
      (
        document.querySelector("#" + input.name) as HTMLInputElement
      ).addEventListener("keyup", (e) => {
        e.preventDefault();
        if (e.key === "Enter") {
          submitBtn.click();
        }
      });
    }

    submitBtn.addEventListener("click", this.props.submitFn);
  }
}
