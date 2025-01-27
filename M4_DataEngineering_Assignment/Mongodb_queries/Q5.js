db.warehouses.aggregate([
    {
      $geoNear: {
        near: { type: "Point", coordinates: [-74.006, 40.7128] },
        distanceField: "distance",
        maxDistance: 50000, // 50 kilometers
        spherical: true
      }
    },
    {
      $match: {
        products: "P001"
      }
    },
    {
      $project: {
        warehouseId: 1,
        location: 1,
        distance: 1
      }
    }
  ])
  