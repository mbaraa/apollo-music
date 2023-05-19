import { useNavigate } from "react-router-dom";

export default function Mobile() {
  const navigate = useNavigate();
  return (
    <div className="w-[100vw] h-[100vh] bg-[#273D50]">
      <img className="pt-[22px] ml-[-102px]" src="/apollo-music-logo.png" />
      <div className="text-white pl-[20px] pt-[60px]">
        <h1 className="text-[48px]">Apollo Music</h1>
        <h2 className="pt-[20px] text-[24px]">
          The Coolest Cloud Music Player!
        </h2>
      </div>
      <div className="absolute bottom-[40px] left-[50%] translate-x-[-50%]">
        <button
          className="bg-white text-black rounded-[20px] p-[13px] w-[228px] h-[54px] flex justify-between items-center"
          onClick={() => navigate("/intro")}
        >
          <span className="text-[17px]">Discover the Beat</span>{" "}
          <img className="inline" src="/headset-icon.png" />
        </button>
      </div>
    </div>
  );
}
