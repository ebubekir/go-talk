// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
import {getAuth} from "@firebase/auth";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
    apiKey: "AIzaSyCKE_wGegIF0GwL0LzDH0SxbzsEBn6dHi8",
    authDomain: "gotalk-85909.firebaseapp.com",
    projectId: "gotalk-85909",
    storageBucket: "gotalk-85909.firebasestorage.app",
    messagingSenderId: "11125079224",
    appId: "1:11125079224:web:bb6469059cfd9f713ec84d",
    measurementId: "G-ZH02DMTSDF"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);
