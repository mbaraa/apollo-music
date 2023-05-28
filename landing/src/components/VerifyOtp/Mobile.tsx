import { FormEvent, useState } from "react";
import ResendEmail from "../../assets/ResendEmail.tsx";
import SelectPlan from "../SelectPlan/index.tsx";

export default function Mobile() {
  const [verificationCode, setVerificationCode] = useState("");
  const [showSelectPlan, setShowSelectPlan] = useState(false);

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
      .then(async (resp) => {
        const respBody = await resp.json();
        if (!resp.ok) {
          window.alert(respBody["errorMsg"]);
          return;
        }
        localStorage.setItem("checkoutToken", respBody["data"]["token"]);
        localStorage.removeItem("otpToken");
        setShowSelectPlan(true);
      })
      .catch((err) => {
        window.alert(err);
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
      {showSelectPlan ? (
        <SelectPlan />
      ) : (
        <div className="w-[100vw] h-[100vh] bg-dark-primary text-dark-secondary">
          <h1 className="text-[30px] pt-[32px] ml-[30px]">Verify your email</h1>
          <h2 className="text-[20px] mt-[20px] ml-[30px]">
            Check your inbox for an OTP from us to unlock your music library.
          </h2>
          <div className="absolute left-[50%] translate-x-[-50%] w-[330px]">
            <form onSubmit={handleSubmit} className="pt-[44px]">
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
                className="bg-dark-accent w-[50px] border-[2px] border-dark-accent rounded-r-[20px] p-[12px] font-[Comfortaa] text-center text-[20px] bg-[url('/verify-otp.svg')] bg-no-repeat bg-center cursor-pointer hover:opacity-[0.9]"
              />
              <button
                type="reset"
                className="bg-dark-accent ml-[10px] text-dark-neutral rounded-[20px] w-[70px] relative top-[-2px] h-[60px] text-[16px] inline-block"
                title="Resend OTP"
                onClick={resendOtp}
              >
                <ResendEmail color="#051220" width="40px" height="40px" />
              </button>
            </form>

            <div className="pt-[20px]">
              <img src="/email-image.png" />
            </div>
          </div>
        </div>
      )}
    </>
  );
}
