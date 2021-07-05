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
    const submitBtn = document.querySelector("#submit") as HTMLInputElement;
    const inputs = this.props.inputs.map((input) => {
      return (
        <input
          type={input.type}
          id={input.name}
          key={input.name}
          name={input.name}
          placeholder={capitalizeFirstLetter(input.name)}
          required={input.required}
          minLength={input.minLength}
          onChange={input.onChangeFn}
          onKeyUp={(e) => {
            e.preventDefault();
            if (e.key === "Enter" && !submitBtn.disabled) {
              submitBtn.click();
            }
          }}
        />
      );
    });

    return (
      <div className="userForm">
        <h1>{capitalizeFirstLetter(this.props.title)}</h1>

        <form
          onSubmit={(e) => {
            e.preventDefault();
            this.props.submitFn();
          }}
        >
          {inputs}

          <input
            id="submit"
            type="submit"
            value={this.props.title}
            disabled={true}
          />
        </form>
      </div>
    );
  }

  componentDidMount() {
    setTitle(capitalizeFirstLetter(this.props.title));
  }
}
