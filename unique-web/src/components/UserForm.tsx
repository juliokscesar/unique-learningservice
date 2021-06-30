import React from "react";
import { toggleElementById, validateEmail } from "../utils";

type UserFormSubmitFunc = (email: string, password: string) => void;

type UserFormProps = {
  title: string;
  submitFn: UserFormSubmitFunc;
};

export const UserForm = (props: UserFormProps) => {
  return (
    <div className="userForm">
      <form action="/">
        <h3>{ props.title }</h3>

        <label>Email:
          <input
            id="email"
            name="email"
            type="text"
            placeholder="Email"
            required={ true }
            onChange={ (e) => {
              const pass = document.querySelector("#password") as HTMLInputElement;
              toggleElementById(
                "userSubmit",
                (validateEmail(e.target.value) && pass.value.length >= 8)
              );
            }}
          />
        </label>
  
        <br />

        <label>Password:
          <input
            id="password"
            name="password"
            type="password"
            placeholder="Password"
            minLength={ 8 }
            required={ true }
            onChange={ (e) => {
              const email = document.querySelector("#email") as HTMLInputElement;
              toggleElementById(
                "userSubmit",
                (e.target.value.length >= 8 && validateEmail(email.value))
              );
            }}
          />
        </label>

        <input 
          id="userSubmit"
          type="submit"
          value={ props.title }
          disabled={true}
          onSubmit={ () => {
            const email = document.querySelector("#email") as HTMLInputElement;
            const password = document.querySelector("#password") as HTMLInputElement;
            props.submitFn(email.value, password.value);
          }}
        />

      </form>
    </div>
  )
}
