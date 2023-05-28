import Desktop from "./Desktop";
import Mobile from "./Mobile";

export default function VerifyOtp() {
  return (
    <div>
      <div className="md:hidden">
        <Mobile />
      </div>
      <div className="hidden md:block">
        <Desktop />
      </div>
    </div>
  );
}
