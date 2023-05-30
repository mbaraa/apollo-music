import { FormEvent, useState } from "react";
import VerifyOtp from "../VerifyOtp";

export default function Mobile() {
  const [user, setUser] = useState({ fullName: "", email: "", password: "" });
  const [showOtp, setShowOtp] = useState(false);

  const signupUser = async () => {
    await fetch(`${import.meta.env.VITE_BACKEND_ADDRESS}/auth/signup/email`, {
      method: "POST",
      mode: "cors",
      body: JSON.stringify(user),
    })
      .then(async (resp) => {
        const respBody = await resp.json();
        if (!resp.ok) {
          window.alert(respBody["errorMsg"]);
          return;
        }
        localStorage.setItem("otpToken", respBody["data"]["token"]);
        setShowOtp(true);
      })
      .catch((err) => {
        window.alert(err);
      });
  };

  const handleSignup = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    signupUser();
  };

  return (
    <>
      {showOtp ? (
        <VerifyOtp />
      ) : (
        <div className="w-[100vw] h-[100vh] bg-dark-primary text-dark-secondary">
          <div className="flex justify-between w-[100vw]">
            <h1 className="text-white text-[28px] pt-[32px] pl-[30px]">
              Unlock a world of YOUR music with us!
            </h1>
            <img
              width={248}
              height={218}
              className="mt-[-20px] w-[124px] h-[218px] object-cover object-left"
              src="/cassette.png"
            />
          </div>

          <div className="absolute left-[50%] translate-x-[-50%]">
            <form
              onSubmit={handleSignup}
              className="pt-[20px] font-IBMPlexSans"
            >
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
                type="submit"
                className="bg-dark-accent text-dark-neutral w-[330px] h-[48px] rounded-[20px] mt-[25px] text-[24px] cursor-pointer"
                value="Sign up"
              />
            </form>

            <h3 className="w-full text-center mt-[35px]">
              Already have an account? Log in{" "}
              <a
                href="https://apollo-music.app/sign-in"
                className="underline hover:opacity-[0.9]"
              >
                here
              </a>{" "}
            </h3>
          </div>
        </div>
      )}
    </>
  );
}
