import Header from "../components/Header";
import Subheader from "../components/Subheading";
import InputBoxes from "../components/InputBoxes";
import AuthBottom from "../components/AuthBottom";

type SignupProps = {
  onClose: () => void;
};

const SignupPage = ({ onClose }: SignupProps) => {
  return (
    <div className="fixed inset-0 bg-black/40 backdrop-blur-sm flex items-center justify-center z-50">

      <div className="bg-white w-105 rounded-2xl shadow-2xl p-8 relative flex flex-col gap-5">

        <button
          onClick={onClose}
          className="absolute top-4 right-4 text-gray-500 hover:text-black text-lg font-bold"
        >
          ✕
        </button>

        <Header text="Welcome to Wintrest" />
        <Subheader text="Find new ideas to try" />

        <div className="flex flex-col gap-4">
          <InputBoxes heading="Email" placeholder="Enter your email" />
          <InputBoxes heading="Password" placeholder="Create a password" />
        </div>

        <div className="flex items-center gap-2 text-sm text-gray-600 hover:text-black cursor-pointer">
          Password Tips
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" fill="currentColor" className="size-4">
          <path fillRule="evenodd" d="M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0ZM9 5a1 1 0 1 1-2 0 1 1 0 0 1 2 0ZM6.75 8a.75.75 0 0 0 0 1.5h.75v1.75a.75.75 0 0 0 1.5 0v-2.5A.75.75 0 0 0 8.25 8h-1.5Z" clipRule="evenodd" />
          </svg>
        </div>

        <button className="bg-red-500 hover:bg-red-600 text-white font-semibold py-3 rounded-xl transition">
          Continue
        </button>
        <AuthBottom
          text="Already a member?"
          linktext="Log in"
          onClick={onClose}
        />

      </div>
    </div>
  );
};

export default SignupPage;