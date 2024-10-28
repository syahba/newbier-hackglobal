function CardTransaction({ data }) {
  return (
    <div className="flex items-center rounded p-3 shadow-md">
      <div className="flex flex-grow flex-col gap-1">
        <h5 className="text-sm font-bold">{data.name}</h5>
        <p>{data.product}</p>
      </div>
      <p className="text-sm">{data.price}</p>
    </div>
  );
}

export default CardTransaction;
