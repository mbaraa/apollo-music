import { useNavigate } from "react-router-dom";

export default function Mobile() {
  const navigate = useNavigate();
  return (
    <div className="w-[100vw] h-[100vh] bg-[#273D50]">
      <img
        className="pt-[92px] relative translate-x-[-50%] left-[50%]"
        src="/listing-to-music.png"
      />
      <div className="text-white pl-[20px] pt-[42px]">
        <h2 className="pt-[20px] text-[24px]">
          Enjoy ad-free music, or self hosted music with Apollo Music!
        </h2>
      </div>
      <div className="absolute bottom-[40px] left-[50%] translate-x-[-50%]">
        <div className="relative w-[210px] font-[700] text-[16px/20px]">
          <button
            className="inline bg-white text-black rounded-[20px] p-[13px] w-[110px] h-[47px] text-center"
            onClick={() => navigate("/sign-up")}
          >
            <span className="">Sign up</span>
          </button>
          <button
            className="inline bg-[#add8fb] text-black rounded-[20px] ml-[-18px] p-[13px] w-[110px] h-[47px] shadow-md"
            onClick={() =>
              window.open("https://apollo-music.app/sign-in", "_self")
            }
          >
            <span>Log in </span>
            <span>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth={2.5}
                stroke="currentColor"
                className="w-5 h-5 inline font-bold"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M13.5 4.5L21 12m0 0l-7.5 7.5M21 12H3"
                />
              </svg>
            </span>
          </button>
        </div>
      </div>
    </div>
  );
}
