import axios from 'axios';

const API_URL = 'http://localhost:8080/api/';

class AuthService {
	login(user) {
		const params = new URLSearchParams();
		params.append('username', user.username);
		params.append('password', user.password);
		return axios
		.post(API_URL + 'signin', params)
		.then(response => {
			console.log(response.data.status)
			if (response.data.status == "success") {
			// if (response.data.accessToken) {
				localStorage.setItem('user', JSON.stringify(response.data));
			}

			return response.data;
		});
	}

	logout() {
		localStorage.removeItem('user');
	}

	// register(user) {
	// 	return axios.post(API_URL + 'signup', {
	// 		username: user.username,
	// 		email: user.email,
	// 		password: user.password
	// 	});
	// }
}

export default new AuthService();