import { QueryClient, QueryClientProvider } from "react-query";
import { AuthProvider } from "./contexts/AuthContext";
import { UserProvider } from "./contexts/UserContext";
import { Routes, Route } from "react-router-dom";
import { NavBar } from "./components/navbar/NavBar"
import { Login } from "./pages/login/Login";
import { Signup } from "./pages/signup/Signup";
import { ForgotPassword } from "./pages/forgotPassword/ForgotPassword";
import { PrivateRoute } from "./components/privateRoute/PrivateRoute";
import { Profile } from "./pages/profile/Profile";
import { UpdateProfile } from "./pages/updateProfile/UpdateProfile"
import { Courses } from "./pages/courses/Courses"
import { Course } from "./pages/course/Course";
import { Mentees } from "./pages/mentees/Mentees";
import { Mentee } from './pages/mentee/Mentee';
import { Admin } from "./pages/admin/Admin";
import { AssignCourse } from "./pages/assignCourse/AssignCourse";

const queryClient = new QueryClient();

const App = () => {
  return (
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <UserProvider>
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
            <Route path='/mentees' element={
              <PrivateRoute>
                <Mentees />
              </PrivateRoute>}
            />
            <Route path='/mentor/:mentee/:courseId' element={
              <PrivateRoute>
                <Mentee />
              </PrivateRoute>}
            />
            <Route path='/admin' element={
              <PrivateRoute>
                <Admin />
              </PrivateRoute>}
            />
            <Route path='/admin/assign' element={
              <PrivateRoute>
                <AssignCourse />
              </PrivateRoute>}
            />
          </Routes>
        </UserProvider>
      </AuthProvider>
    </QueryClientProvider>
  )
}

export default App