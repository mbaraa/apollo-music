import { FormEvent, useState } from "react";
import { useNavigate } from "react-router-dom";

export default function OTP() {
  const navigate = useNavigate();
  const [verificationCode, setVerificationCode] = useState("");

  const verifyOtp = async () => {
    await fetch(`${import.meta.env.VITE_BACKEND_ADDRESS}/auth/otp/verify`, {
      method: "POST",
      mode: "cors",
      body: JSON.stringify({
        verificationCode: verificationCode,
      }),
      headers: {
        Authorization: localStorage.getItem("otpToken") ?? "",
      },
    })
      .then((resp) => {
        if (resp.ok) {
          localStorage.removeItem("otpToken");
        }
        return resp.json();
      })
      .then((resp) => {
        localStorage.setItem("checkoutToken", resp["data"]["token"]);
        navigate("/checkout");
      })
      .catch((err) => {
        console.error(err);
      });
  };

  const resendOtp = async () => {
    await fetch(`${import.meta.env.VITE_BACKEND_ADDRESS}/auth/otp/resend`, {
      method: "GET",
      headers: {
        Authorization: localStorage.getItem("otpToken") ?? "",
      },
    }).catch((err) => {
      console.error(err);
    });
  };

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    verifyOtp();
  };

  return (
    <div className="w-[100vw] h-[100vh] bg-dark-primary text-dark-secondary">
      <h1 className="text-[30px] pt-[32px] ml-[30px]">Verify your email</h1>
      <h2 className="text-[20px] mt-[20px] ml-[30px]">
        Check your inbox for an OTP from us to unlock your music library.
      </h2>
      <form onSubmit={handleSubmit} className="pt-[44px] ml-[30px]">
        <input
          type="text"
          maxLength={6}
          placeholder="XXXXXX"
          required
          value={verificationCode}
          onChange={(e) => {
            setVerificationCode(e.target.value);
          }}
          className="bg-dark-primary border-[2px] border-dark-accent rounded-l-[20px] p-[12px] w-[200px] font-[Comfortaa] text-center text-[20px]"
        />
        <input
          type="submit"
          value=" "
          className="bg-dark-accent border-[2px] border-dark-accent rounded-r-[20px] p-[12px] font-[Comfortaa] text-center text-[20px] bg-[url('/verify-otp.png')] bg-no-repeat bg-center w-[50px] cursor-pointer hover:opacity-[0.9]"
        />
      </form>
      <button
        className="bg-dark-accent text-dark-neutral mt-[10px] ml-[30px] rounded-[20px] p-[13px] w-[130px] h-[34px] flex items-center text-center text-[16px]"
        onClick={resendOtp}
      >
        Resend OTP
      </button>
      <div className="">
        <img src="/email-image.png" />
      </div>
    </div>
  );
}
