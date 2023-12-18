import { createContext, useContext, useState, useEffect } from "react";
import { createUserWithEmailAndPassword, signInWithEmailAndPassword, sendPasswordResetEmail, onAuthStateChanged, signOut, updateEmail, updatePassword, updateProfile } from "firebase/auth";
import { auth } from "../firebase";

const AuthContext = createContext();

export const useAuth = () => {
  return useContext(AuthContext);
}

export const AuthProvider = ({ children }) => {
  const [currentUser, setCurrentUser] = useState();
  const [loading, setLoading] = useState(true);

  const signup = (email, password) => {
    return createUserWithEmailAndPassword(auth, email, password);
  }

  const displayName = (name) => {
    return updateProfile(auth.currentUser, {
      displayName: name,
    })
  }

  const login = (email, password) => {
    return signInWithEmailAndPassword(auth, email, password);
  }

  const logout = () => {
    return signOut(auth);
  }

  const resetPassword = (email) => {
    return sendPasswordResetEmail(auth, email);
  }

  const emailUpdate = (email) => {
    return updateEmail(auth.currentUser, email);
  }

  const passwordUpdate = async (password) => {
    try {
      await updatePassword(auth.currentUser, password);
      console.log("Password updated successfully");
    } catch (error) {
      console.log("Password update failed");
    }
  }

  useEffect(() => {
    const unsubscribe = onAuthStateChanged(auth, (user) => {
      setCurrentUser(user);
      setLoading(false);
    })

    return unsubscribe
  }, [])

  const value = {
    currentUser,
    login,
    signup,
    displayName,
    logout,
    resetPassword,
    emailUpdate,
    passwordUpdate
  }

  return (
    <AuthContext.Provider value={value}>
      {!loading && children}
    </AuthContext.Provider>
  )
}
