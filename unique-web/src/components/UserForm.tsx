import React from "react";
import { capitalizeFirstLetter } from "../utils";

type FnEventChange = (e: React.ChangeEvent<HTMLInputElement>) => void;
type FnSubmit = () => void;

type UserFormProps = {
  title: string;
  inputs: {
      name: string,
      type: string,
      required: boolean,
      minLength?: number,
      onChangeFn?: FnEventChange
    }[];
    submitFn: FnSubmit;
}

export class UserForm extends React.Component<UserFormProps> {
  render() {
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
        />
      );
    });

    return (
      <div className="userForm">
        <h1>{capitalizeFirstLetter(this.props.title)}</h1>
        
        {inputs}
        
        <button id="submit" disabled={true}>{capitalizeFirstLetter(this.props.title)}</button>

      </div>
    );
  }

  componentDidMount() {
    (document.querySelector("#submit") as HTMLInputElement).addEventListener("click", this.props.submitFn);
  }
}
