import Mobile from "./Mobile";

export default function Landing() {
  return (
    <div>
      <div className="md:hidden">
        <Mobile />
      </div>
      <div className="hidden md:block">desktop</div>
    </div>
  );
}
