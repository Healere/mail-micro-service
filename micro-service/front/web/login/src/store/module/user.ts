import storageService from "../../service/storageService";
import userService from "../../service/userService";

const userModule = {
    namespaced: true,
    state() {
      return {
        token: storageService.get(storageService.USER_TOKEN),
        userInfo: storageService.get(storageService.USER_INFO) ? JSON.parse(storageService.get(storageService.USER_INFO)) : null,
      };
    },

    mutations: {
      SET_TOKEN(state, token) {
        //更新本地缓存
        storageService.set(storageService.USER_TOKEN, token);
        // 更新state
        state.token = token;
      },
      SET_USERINFO(state, userInfo) {
        //更新本地缓存
        storageService.set(storageService.USER_INFO, JSON.stringify(userInfo));
        // 更新state
        state.userInfo = userInfo;
      },
    },
    actions: {
      register(context, {name,email,password}) {
        return new Promise((resolve, reject) => {
          userService.register({name,email,password})
            .then((res) => {
              //保存token
              //this.$store.commit('userModule/SET_TOKEN',res.data.data.token);
              //this.$store.commit('userModule/TEST',{});
              context.commit("SET_TOKEN", res.data.data.token);
              //localStorage.setItem(storageService.USER_TOKEN,res.data.data.token)
              return userService.info();
            })
            .then((res) => {
              // 保存用户信息
              //this.$store.commit('userModule/SET_USERINFO',response.data.data.user);
              context.commit("SET_USERINFO", res.data.data.user);
              // storageService.set(
              //   storageService.USER_INFO,
              //   JSON.stringify(response.data.data.user)
              // );
              resolve(res);
            })
            .catch((err) => {
              reject(err);
            });
        });
      },
      login(context, {name,email,password}) {
        return new Promise((resolve, reject) => {
          userService.login({name,email,password})
            .then((res) => {
              context.commit("SET_TOKEN", res.data.data.token);
              return userService.info();
            })
            .then((res) => {
              context.commit("SET_USERINFO", res.data.data.user);
              resolve(res);
            })
            .catch((err) => {
              reject(err);
            });
        });
      },
      logout({commit}) {
        // 清除token
        commit("SET_TOKEN", "");
        storageService.set(storageService.USER_TOKEN, "");
        //清除用户信息
        commit("SET_USERINFO", "");
        storageService.set(storageService.USER_INFO, "");
      },
    }
  };

export default userModule;
