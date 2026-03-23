import type { JSX } from "react";

const SearchBar = (): JSX.Element => {
  return (
    <div className="relative flex items-center">
      <div className="absolute left-3 flex items-center pointer-events-none">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" className="size-4 text-gray-400">
          <path fillRule="evenodd" d="M9.965 11.026a5 5 0 1 1 1.06-1.06l2.755 2.754a.75.75 0 1 1-1.06 1.06l-2.755-2.754ZM10.5 7a3.5 3.5 0 1 1-7 0 3.5 3.5 0 0 1 7 0Z" clipRule="evenodd" />
        </svg>
      </div>
      <input
        className="h-10 w-105 rounded-2xl bg-gray-100 hover:bg-gray-200 focus:ring-2 focus:ring-blue-300 pl-9 pr-4 outline-none"
        placeholder="Search"
      />
    </div>
  );
};

export default SearchBar;