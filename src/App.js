import { QueryClient, QueryClientProvider } from "react-query";
import { AuthProvider } from "./contexts/AuthContext";
import { Routes, Route } from "react-router-dom";
import { NavBar } from "./components/navbar/NavBar"
import { Login } from "./components/auth/Login";
import { Signup } from "./components/auth/Signup";
import { ForgotPassword } from "./components/auth/ForgotPassword";
import { PrivateRoute } from "./components/auth/PrivateRoute"
import { Profile } from "./components/auth/Profile";
import { UpdateProfile } from "./components/auth/UpdateProfile"
import { Courses } from "./components/courses/Courses";
import { Course } from "./components/course/Course";
import { Mentor } from "./components/mentor/Mentor";
import "./App.css";

const queryClient = new QueryClient();

const App = () => {
  return (
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <NavBar />
        <Routes>
          <Route path="/" element={<Login />} />
          <Route path="/signup" element={<Signup />} />
          <Route path="/login" element={<Login />} />
          <Route path='/forgot-password' element={<ForgotPassword />} />
          <Route path='/profile' element={
            <PrivateRoute>
              <Profile />
            </PrivateRoute>}
          />
          <Route path='/update-profile' element={
            <PrivateRoute>
              <UpdateProfile />
            </PrivateRoute>}
          />
          <Route path='/courses' element={
            <PrivateRoute>
              <Courses />
            </PrivateRoute>}
          />
          <Route path='/courses/:courseId' element={
            <PrivateRoute>
              <Course />
            </PrivateRoute>}
          />
          <Route path='/mentor' element={
            <PrivateRoute>
              <Mentor />
            </PrivateRoute>}
          />
        </Routes>
      </AuthProvider>
    </QueryClientProvider>
  )
}

export default App