import axios from "axios";
import localStorage from '../util/localStorage';
import _ from 'lodash';

const myAxios = axios.create({
  baseURL: process.env.NODE_ENV === 'production' ? 'https://uno-fievsqoyiq-uw.a.run.app' : 'http://localhost:8080'
});

//myAxios.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';

myAxios.interceptors.request.use(function (config) {
  const token = localStorage.get('token');
  if (token) {
    config.headers.Authorization = `Token ${token}`;
  }

  return config;
}, function (error) {
  // Do something with request error
  return Promise.reject(error);
});

myAxios.interceptors.response.use(function (response) {
  // Any status code that lie within the range of 2xx cause this function to trigger
  const userMessage = _.get(response, 'data.userMessage', false);
  if (userMessage) {
    // TODO: Once we have a snack service
    // store.dispatch('addMessage', userMessage);
  }
  return response;
}, function (error) {
  // Any status codes that falls outside the range of 2xx cause this function to trigger
  const userMessage = _.get(error, 'response.data', false);
  if (userMessage) {
    // TODO: Once we have a snack service
    // store.dispatch('showError', userMessage);
  }

  return Promise.reject(error);
});

export default myAxios;

