/* tslint:disable */
/* eslint-disable */
/**
 * First social
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import { Configuration } from './configuration';
import globalAxios, { AxiosPromise, AxiosInstance, AxiosRequestConfig } from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from './common';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, RequestArgs, BaseAPI, RequiredError } from './base';

/**
 * 
 * @export
 * @interface ErrorResponse
 */
export interface ErrorResponse {
    /**
     * 
     * @type {Array<ErrorResponseErrors>}
     * @memberof ErrorResponse
     */
    'errors'?: Array<ErrorResponseErrors>;
}
/**
 * 
 * @export
 * @interface ErrorResponseErrors
 */
export interface ErrorResponseErrors {
    /**
     * 
     * @type {string}
     * @memberof ErrorResponseErrors
     */
    'message'?: string;
    /**
     * 
     * @type {object}
     * @memberof ErrorResponseErrors
     */
    'params'?: object;
}
/**
 * 
 * @export
 * @interface LoginCredentials
 */
export interface LoginCredentials {
    /**
     * 
     * @type {string}
     * @memberof LoginCredentials
     */
    'email': string;
    /**
     * 
     * @type {string}
     * @memberof LoginCredentials
     */
    'password': string;
}
/**
 * 
 * @export
 * @interface Profile
 */
export interface Profile {
    /**
     * 
     * @type {number}
     * @memberof Profile
     */
    'id'?: number;
    /**
     * 
     * @type {string}
     * @memberof Profile
     */
    'email': string;
    /**
     * 
     * @type {string}
     * @memberof Profile
     */
    'password': string;
    /**
     * 
     * @type {string}
     * @memberof Profile
     */
    'firstName': string;
    /**
     * 
     * @type {string}
     * @memberof Profile
     */
    'lastName': string;
    /**
     * 
     * @type {number}
     * @memberof Profile
     */
    'age': number;
    /**
     * 
     * @type {string}
     * @memberof Profile
     */
    'city': string;
    /**
     * 
     * @type {string}
     * @memberof Profile
     */
    'interests': string;
}

/**
 * PrivateApi - axios parameter creator
 * @export
 */
export const PrivateApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Get current profile
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getCurrentProfile: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/profile`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication jwt required
            // http bearer authentication required
            await setBearerAuthToObject(localVarHeaderParameter, configuration)


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Get profile by id
         * @param {number} id ID of profile to return
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getProfile: async (id: number, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'id' is not null or undefined
            assertParamExists('getProfile', 'id', id)
            const localVarPath = `/profiles/{id}`
                .replace(`{${"id"}}`, encodeURIComponent(String(id)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication jwt required
            // http bearer authentication required
            await setBearerAuthToObject(localVarHeaderParameter, configuration)


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary List profiles
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        listProfiles: async (options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/profiles`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication jwt required
            // http bearer authentication required
            await setBearerAuthToObject(localVarHeaderParameter, configuration)


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * PrivateApi - functional programming interface
 * @export
 */
export const PrivateApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = PrivateApiAxiosParamCreator(configuration)
    return {
        /**
         * 
         * @summary Get current profile
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getCurrentProfile(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Profile>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getCurrentProfile(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Get profile by id
         * @param {number} id ID of profile to return
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async getProfile(id: number, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Profile>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.getProfile(id, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary List profiles
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async listProfiles(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<Profile>>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.listProfiles(options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
    }
};

/**
 * PrivateApi - factory interface
 * @export
 */
export const PrivateApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = PrivateApiFp(configuration)
    return {
        /**
         * 
         * @summary Get current profile
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getCurrentProfile(options?: any): AxiosPromise<Profile> {
            return localVarFp.getCurrentProfile(options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Get profile by id
         * @param {number} id ID of profile to return
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        getProfile(id: number, options?: any): AxiosPromise<Profile> {
            return localVarFp.getProfile(id, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary List profiles
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        listProfiles(options?: any): AxiosPromise<Array<Profile>> {
            return localVarFp.listProfiles(options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * PrivateApi - object-oriented interface
 * @export
 * @class PrivateApi
 * @extends {BaseAPI}
 */
export class PrivateApi extends BaseAPI {
    /**
     * 
     * @summary Get current profile
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PrivateApi
     */
    public getCurrentProfile(options?: AxiosRequestConfig) {
        return PrivateApiFp(this.configuration).getCurrentProfile(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Get profile by id
     * @param {number} id ID of profile to return
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PrivateApi
     */
    public getProfile(id: number, options?: AxiosRequestConfig) {
        return PrivateApiFp(this.configuration).getProfile(id, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary List profiles
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PrivateApi
     */
    public listProfiles(options?: AxiosRequestConfig) {
        return PrivateApiFp(this.configuration).listProfiles(options).then((request) => request(this.axios, this.basePath));
    }
}


/**
 * PublicApi - axios parameter creator
 * @export
 */
export const PublicApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Register new profile
         * @param {Profile} profile 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createProfile: async (profile: Profile, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'profile' is not null or undefined
            assertParamExists('createProfile', 'profile', profile)
            const localVarPath = `/register`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(profile, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Logs user into the system
         * @param {LoginCredentials} [loginCredentials] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        login: async (loginCredentials?: LoginCredentials, options: AxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/login`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(loginCredentials, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * PublicApi - functional programming interface
 * @export
 */
export const PublicApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = PublicApiAxiosParamCreator(configuration)
    return {
        /**
         * 
         * @summary Register new profile
         * @param {Profile} profile 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async createProfile(profile: Profile, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<void>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.createProfile(profile, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
        /**
         * 
         * @summary Logs user into the system
         * @param {LoginCredentials} [loginCredentials] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async login(loginCredentials?: LoginCredentials, options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<void>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.login(loginCredentials, options);
            return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
        },
    }
};

/**
 * PublicApi - factory interface
 * @export
 */
export const PublicApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = PublicApiFp(configuration)
    return {
        /**
         * 
         * @summary Register new profile
         * @param {Profile} profile 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        createProfile(profile: Profile, options?: any): AxiosPromise<void> {
            return localVarFp.createProfile(profile, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Logs user into the system
         * @param {LoginCredentials} [loginCredentials] 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        login(loginCredentials?: LoginCredentials, options?: any): AxiosPromise<void> {
            return localVarFp.login(loginCredentials, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * PublicApi - object-oriented interface
 * @export
 * @class PublicApi
 * @extends {BaseAPI}
 */
export class PublicApi extends BaseAPI {
    /**
     * 
     * @summary Register new profile
     * @param {Profile} profile 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PublicApi
     */
    public createProfile(profile: Profile, options?: AxiosRequestConfig) {
        return PublicApiFp(this.configuration).createProfile(profile, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Logs user into the system
     * @param {LoginCredentials} [loginCredentials] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof PublicApi
     */
    public login(loginCredentials?: LoginCredentials, options?: AxiosRequestConfig) {
        return PublicApiFp(this.configuration).login(loginCredentials, options).then((request) => request(this.axios, this.basePath));
    }
}


