# CubeIt Planning Notes

## What we want?

- Users

For a first step, I want users to be able to log in through their WCA account.

- [x] [Check WCA API docs for oauth](https://www.worldcubeassociation.org/help/api)

I've created a WCA application for the authentication, which can be checked here: [WCA App](https://www.worldcubeassociation.org/oauth/applications/1317)

`https://www.worldcubeassociation.org/api/v0/me` sends back the following JSON response:

```json
{
  "me": {
    "id": 11004,
    "created_at": "2016-02-06T11:42:15.000Z",
    "updated_at": "2025-04-25T05:30:49.000Z",
    "name": "Dániel Lázár",
    "wca_id": "2014LAZA04",
    "gender": "m",
    "country_iso2": "HU",
    "url": "https://www.worldcubeassociation.org/persons/2014LAZA04",
    "country": {
      "id": "Hungary",
      "name": "Hungary",
      "continent_id": "_Europe",
      "iso2": "HU"
    },
    "delegate_status": null,
    "class": "user",
    "teams": [],
    "avatar": {
      "id": 66689,
      "status": "approved",
      "thumbnail_crop_x": 155,
      "thumbnail_crop_y": 0,
      "thumbnail_crop_w": 1863,
      "thumbnail_crop_h": 1863,
      "url": "https://avatars.worldcubeassociation.org/qs67jk4b477nafspfk5rgjsulhm3",
      "thumb_url": "https://avatars.worldcubeassociation.org/kapxualu4z4x7ncjtobcv6ecdm5x",
      "is_default": false,
      "can_edit_thumbnail": true
    }
  }
}
```

What if I want to fetch other existing users?

## DataBase, Tables, etc.
