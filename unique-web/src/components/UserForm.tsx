import React from "react";

type UserFormProps = {
  title: string;
};

export const UserForm = ({ title }: UserFormProps) => {
  return (
    <div className="userForm">
      <form action="/">
        <h3>{ title }</h3>

        <label htmlFor="email">Email:</label>
        <input id="email" name="email" type="text" placeholder="Email" />
  
        <br />

        <label htmlFor="password">Password:</label>
        <input id="password" name="password" type="password" placeholder="Password" />

        <input type="submit" value={ title } />
      </form>
    </div>
  )
}
