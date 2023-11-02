import { useQuery } from 'react-query';
import axios from 'axios';
import { QueryClient, QueryClientProvider } from "react-query";
import { AuthProvider, useAuth } from "./contexts/AuthContext";
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
import { UserProvider } from "./contexts/UserContext";
import { Admin } from "./components/admin/Admin";

const queryClient = new QueryClient();

const App = () => {
  // const { data: userDetails, isLoading, isError } = useQuery(["userDetails"], async () => {
  //   const response = await axios.get(`http://localhost:4000/userDetails`);
  //   return response.data;
  // });
  // const { currentUser } = useAuth();

  // if (isLoading) {
  //   return <h3>Loading...</h3>
  // }

  // if (isError) {
  //   return <h3>Loading...</h3>
  // }
  // const user = currentUser && userDetails?.filter(user => user.email === currentUser.email)
  // const role = currentUser && user[0]?.role;

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
            <Route path='/mentor' element={
              <PrivateRoute>
                <Mentor />
              </PrivateRoute>}
            />
            <Route path='/admin' element={
              <PrivateRoute>
                <Admin />
              </PrivateRoute>}
            />
          </Routes>
        </UserProvider>
      </AuthProvider>
    </QueryClientProvider>
  )
}

export default App