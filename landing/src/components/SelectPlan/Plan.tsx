interface PlanProps {
  title: string;
  price: string;
  size: string;
}

export default function Plan({ title, price, size }: PlanProps) {
  return (
    <div className="p-[20px] w-[300px] h-[200px] border-[2px] border-dark-accent bg-dark-accent2 rounded-[20px] grid grid-cols-1 place-items-center">
      <h1 className="text-[30px] block">{title}</h1>
      <h2 className="text-[25px] block">{price}</h2>
      <h3 className="text-[16px] block">{size} Cloud storage for your music</h3>
    </div>
  );
}
