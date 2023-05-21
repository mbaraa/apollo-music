import { FormEvent, useState } from "react";
import { useNavigate } from "react-router-dom";

export default function Mobile() {
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
    <div className="w-[100vw] h-[100vh] bg-dark-primary text-dark-secondary">
      <div className="flex justify-between">
        <h1 className="text-white text-[28px] pt-[32px] pl-[30px]">
          Unlock a world of YOUR music with us!
        </h1>
        <img className="mr-[-120px] w-[248px] h-[218px]" src="/cassette.png" />
      </div>
      <form onSubmit={handleSignup} className="pt-[0px] pl-[30px]">
        <input
          className="block bg-dark-primary text-dark-secondary border-[2px] border-dark-accent rounded-[20px] p-[16px] h-[56px] w-[330px]"
          type="text"
          placeholder="Full Name"
          required
          value={user.fullName}
          onChange={(e) => {
            user.fullName = e.target.value;
            setUser({ ...user });
          }}
        />
        <input
          className="block bg-dark-primary text-dark-secondary border-[2px] border-dark-accent rounded-[20px] p-[16px] h-[56px] w-[330px] mt-[20px]"
          type="email"
          placeholder="Email"
          required
          value={user.email}
          onChange={(e) => {
            user.email = e.target.value;
            setUser({ ...user });
          }}
        />
        <input
          className="block bg-dark-primary text-dark-secondary border-[2px] border-dark-accent rounded-[20px] p-[16px] h-[56px] w-[330px] mt-[20px]"
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
        <input
          className="block bg-dark-primary text-dark-secondary border-[2px] border-dark-accent rounded-[20px] p-[16px] h-[56px] w-[330px] mt-[20px]"
          type="password"
          minLength={8}
          placeholder="Confirm Password"
          required
          value={confirmPassword}
          onChange={(e) => {
            setConfirmPassword(e.target.value);
          }}
        />
        <input
          type="submit"
          className="bg-dark-accent text-dark-neutral w-[330px] h-[48px] rounded-[20px] mt-[25px] text-[24px] cursor-pointer"
          value="Sign up"
        />
      </form>

      <h3 className="w-full text-center mt-[25px]">
        Already have an account? Log in{" "}
        <a
          href="https://apollo-music.app/sign-in"
          className="underline hover:opacity-[0.9]"
        >
          here
        </a>{" "}
      </h3>
    </div>
  );
}
