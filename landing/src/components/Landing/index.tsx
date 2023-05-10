import { useNavigate } from "react-router-dom";

export default function Landing() {
  const navigate = useNavigate();
  return (
    <>
      Apollo Music, the coolest cloud audio player! <br />{" "}
      <button onClick={() => navigate("/signup")}>Go to signup</button>
    </>
  );
}
