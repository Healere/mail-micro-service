<template>
  <!-- <div id="all"> -->
  <div>
    <div class="allDiv">
      <div class="container right-panel-active">
        <!-- Sign Up -->
        <div class="container__form container--signup">
          <form action="#" class="form" id="form1">
            <h2 class="form__title">Sign Up</h2>
            <input type="text" placeholder="User" class="input" v-model="user.name" />
            <input type="text" placeholder="Email" class="input" v-model="user.email" />
            <input type="password" placeholder="Password" class="input" v-model="user.password" />
            <input class="btn" @click="register(), log()" type="button" value="Sign Up" />
          </form>
        </div>

        <!-- Sign In -->
        <div class="container__form container--signin">
          <form action="#" class="form" id="form2">
            <h2 class="form__title">Sign In</h2>
            <input type="email" placeholder="Email" class="input" v-model="user.email" />
            <input type="password" placeholder="Password" class="input" v-model="user.password" />
            <a href="#" class="link">Forgot your password?</a>
            <input class="btn" @click="login(),log()" type="button" value="Sign In" />
          </form>
        </div>

        <!-- Overlay -->
        <div class="container__overlay">
          <div class="overlay">
            <div class="overlay__panel overlay--left">
              <button class="btn" id="signIn" @click="SignIn">Sign In</button>
            </div>
            <div class="overlay__panel overlay--right">
              <button class="btn" id="signUp" type="button" @click="SignUp">Sign Up</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions } from "vuex";
export default {
  mounted () {
    const signInBtn = document.getElementById("signIn");
    const signUpBtn = document.getElementById("signUp");
    const container = document.querySelector(".container");

    signInBtn.addEventListener("click", () => {
      container.classList.remove("right-panel-active");
    });

    signUpBtn.addEventListener("click", () => {
      container.classList.add("right-panel-active");
    });
  },
  data () {
    return {
      user: {
        name: "",
        email: "",
        password: "",
      },
      logData: "",
    };
  },
  methods: {
    ...mapActions("userModule", { userRegister: 'register' }),
    register () {
      //this.$store.dispatch("userModule/register", this.user).then(() => {
      this.userRegister(this.user).then(() => {
        this.$router.replace({ name: "mail" });
      }).catch((err) =>{
        alert("err:"+err);
      })
    },
    ...mapActions("userModule", { userLogin: 'login' }),
    login () {
      this.userLogin(this.user).then(() => {
        this.$router.replace({ name: "mail" });
      })
    },
    ...mapActions("logModule",{userLog: 'log'}),
    log () {
      this.userLog(this.user,this.logData)
    },
    SignIn () {
      this.$router.replace({ name: 'login' });
    },
    SignUp () {
      this.$router.replace({ name: 'register' });
    },
  },
};
</script>

<style>
@import "../assets/css/login.css";
</style>
