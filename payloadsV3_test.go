package cfclient

const listAppsPayloadV3 = `{
   "pagination": {
      "total_results": 2,
      "total_pages": 2,
      "first": {
         "href": "https://api.example.org/v3/apps?page=1&per_page=1"
      },
      "last": {
         "href": "https://api.example.org/v3/apps?page=2&per_page=1"
      },
      "next": {
         "href": "https://api.example.org/v3/apps?page=2&per_page=1"
      },
   	  "previous": null
   },
   "resources": [
      {
         "metadata": {
            "guid": "af15c29a-6bde-4a9b-8cdf-43aa0d4b7e3c",
            "url": "/v2/apps/af15c29a-6bde-4a9b-8cdf-43aa0d4b7e3c",
            "created_at": "2014-10-10T21:03:13+00:00",
            "updated_at": "2014-11-10T14:07:31+00:00"
         },
         "entity": {
            "name": "app-test",
            "production": false,
            "space_guid": "8efd7c5c-d83c-4786-b399-b7bd548839e1",
            "stack_guid": "2c531037-68a2-4e2c-a9e0-71f9d0abf0d4",
            "buildpack": "https://github.com/cloudfoundry/buildpack-go.git",
            "detected_buildpack": "",
            "detected_buildpack_guid": "0d22f6a1-76c5-417f-ac6c-d9d21463ecbc",
            "environment_json": {
               "FOOBAR": "QUX"
            },
            "memory": 256,
            "instances": 1,
            "disk_quota": 1024,
            "state": "STARTED",
            "version": "97ef1272-9eb6-4839-9df1-5ed4f55b5c45",
            "command": null,
            "console": false,
            "debug": null,
            "staging_task_id": "5879c8d06a10491a879734162000def8",
            "package_state": "PENDING",
            "health_check_http_endpoint": null,
            "health_check_type": "port",
            "health_check_timeout": null,
            "staging_failed_reason": null,
            "staging_failed_description": null,
            "diego": true,
            "docker_image": null,
            "package_updated_at": "2014-11-10T14:08:50+00:00",
            "detected_start_command": "app-launching-service-broker",
            "enable_ssh": true,
            "docker_credentials_json": {
               "redacted_message": "[PRIVATE DATA HIDDEN]"
            },
            "ports": [
               8080
            ],
            "space_url": "/v2/spaces/8efd7c5c-d83c-4786-b399-b7bd548839e1",
            "stack_url": "/v2/stacks/2c531037-68a2-4e2c-a9e0-71f9d0abf0d4",
            "events_url": "/v2/apps/af15c29a-6bde-4a9b-8cdf-43aa0d4b7e3c/events",
            "service_bindings_url": "/v2/apps/af15c29a-6bde-4a9b-8cdf-43aa0d4b7e3c/service_bindings",
            "routes_url": "/v2/apps/af15c29a-6bde-4a9b-8cdf-43aa0d4b7e3c/routes"
         }
      }
   ]
}`

const listAppsPayloadPage2V3 = `{
   "pagination": {
      "total_results": 2,
      "total_pages": 2,
      "first": {
         "href": "https://api.example.org/v3/apps?page=1&per_page=1" 
      },
      "last": {
         "href": "https://api.example.org/v3/apps?page=2&per_page=1"
      },
   	  "next": null,
      "previous": {
         "href": "https://api.example.org/v3/apps?page=1&per_page=1"
      }
   },
   "resources": [
      {
         "metadata": {
            "guid": "f9ad202b-76dd-44ec-b7c2-fd2417a561e8",
            "url": "/v2/apps/f9ad202b-76dd-44ec-b7c2-fd2417a561e8",
            "created_at": "2014-10-10T21:03:13+00:00",
            "updated_at": "2014-11-10T14:07:31+00:00"
         },
         "entity": {
            "name": "app-test2",
            "production": false,
            "space_guid": "8efd7c5c-d83c-4786-b399-b7bd548839e1",
            "stack_guid": "2c531037-68a2-4e2c-a9e0-71f9d0abf0d4",
            "buildpack": "https://github.com/cloudfoundry/buildpack-go.git",
            "detected_buildpack": null,
            "environment_json": {
               "FOOBAR": "QUX"
            },
            "memory": 256,
            "instances": 1,
            "disk_quota": 1024,
            "state": "STARTED",
            "version": "97ef1272-9eb6-4839-9df1-5ed4f55b5c45",
            "command": null,
            "console": false,
            "debug": null,
            "staging_task_id": "5879c8d06a10491a879734162000def8",
            "package_state": "PENDING",
            "health_check_timeout": null,
            "staging_failed_reason": null,
            "docker_image": null,
            "package_updated_at": "2014-11-10T14:08:50+00:00",
            "detected_start_command": "app-launching-service-broker",
            "space_url": "/v2/spaces/8efd7c5c-d83c-4786-b399-b7bd548839e1",
            "stack_url": "/v2/stacks/2c531037-68a2-4e2c-a9e0-71f9d0abf0d4",
            "events_url": "/v2/apps/af15c29a-6bde-4a9b-8cdf-43aa0d4b7e3c/events",
            "service_bindings_url": "/v2/apps/af15c29a-6bde-4a9b-8cdf-43aa0d4b7e3c/service_bindings",
            "routes_url": "/v2/apps/af15c29a-6bde-4a9b-8cdf-43aa0d4b7e3c/routes"
         }
      }
   ]
}`
