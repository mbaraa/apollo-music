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
      .then((resp) => resp.json())
      .then((resp) => {
        localStorage.removeItem("otpToken");
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
    <>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          maxLength={6}
          placeholder="OTP"
          required
          value={verificationCode}
          onChange={(e) => {
            setVerificationCode(e.target.value);
          }}
        />
        <br />
        <input type="submit" placeholder="Next" value="Next" />
      </form>
      <br />
      <button onClick={resendOtp}>Resend OTP</button>
    </>
  );
}
