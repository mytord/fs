function authHeaders(token: string) {
  return {
    "Authorization": `Bearer ${token}`,
  };
}

export default authHeaders;
