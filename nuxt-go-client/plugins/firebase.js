import firebase from "firebase/app";
import "firebase/auth"

if (!firebase.apps.length) {
  firebase.initializeApp({
    apiKey: "AIzaSyC507QuHKWoikVoPiVHeyWuyq_fd1IqwgA",
    authDomain: "go-nuxt-c505212.firebaseapp.com",
    databaseURL: "https://go-nuxt-c505212.firebaseio.com",
    projectId: "go-nuxt-c505212",
    storageBucket: "go-nuxt-c505212.appspot.com",
    messagingSenderId: "431856509725",
    appId: "1:431856509725:web:a8616fedd8af3108deb7f4"
  })
}

export default firebase;
