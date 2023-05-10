import { FormEvent, useState } from "react";
import { useNavigate } from "react-router-dom";

export default function Signup() {
  const navigate = useNavigate();
  const [user, setUser] = useState({ fullName: "", email: "", password: "" });
  const [confirmPassword, setConfirmPassword] = useState("");

  const signupUser = async () => {
    if (user.password !== confirmPassword) {
      window.alert("Passwords don't match!");
    }

    await fetch(`${import.meta.env.VITE_BACKEND_ADDRESS}/auth/signup/email`, {
      method: "POST",
      mode: "cors",
      body: JSON.stringify(user),
    })
      .then((resp) => resp.json())
      .then((resp) => {
        localStorage.setItem("otpToken", resp["data"]["token"]);
        navigate("/verify-otp");
      })
      .catch((err) => {
        console.error(err);
      });
  };

  const handleSignup = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log("hello");
    signupUser();
  };

  return (
    <div>
      <form onSubmit={handleSignup}>
        <input
          type="text"
          placeholder="Full name"
          required
          value={user.fullName}
          onChange={(e) => {
            user.fullName = e.target.value;
            setUser({ ...user });
          }}
        />
        <br />
        <input
          type="email"
          placeholder="Email"
          required
          value={user.email}
          onChange={(e) => {
            user.email = e.target.value;
            setUser({ ...user });
          }}
        />
        <br />
        <input
          type="password"
          minLength={8}
          placeholder="Password"
          required
          value={user.password}
          onChange={(e) => {
            user.password = e.target.value;
            setUser({ ...user });
          }}
        />
        <br />
        <input
          type="password"
          minLength={8}
          placeholder="Confirm password"
          required
          value={confirmPassword}
          onChange={(e) => {
            setConfirmPassword(e.target.value);
          }}
        />
        <br />
        <input type="submit" placeholder="Next" value="Next" />
      </form>
    </div>
  );
}
