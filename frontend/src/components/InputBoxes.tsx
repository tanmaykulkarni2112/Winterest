type Elements = {
  heading: string
  placeholder: string
}

const InputBoxes = ({ heading, placeholder }: Elements) => {
  return (
    <div className="flex flex-col gap-1 w-full">

      <label className="text-sm font-semibold text-gray-700">
        {heading}
      </label>

      <input
        placeholder={placeholder}
        className="w-full px-4 py-3 rounded-xl border border-gray-300 focus:outline-none focus:ring-2 focus:ring-red-400 focus:border-red-400 transition"/>
    </div>
  )
}

export default InputBoxes