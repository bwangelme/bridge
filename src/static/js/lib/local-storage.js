function getData(key) {
  const raw = localStorage.getItem(key)
  if( raw === null ) {
    return null
  }

  const data = JSON.parse(raw)
  return data
}

function setData(key, data) {
  const raw = JSON.stringify(data)
  localStorage.setItem(key, raw)
}

function getUnitMap() {
  const s = 1000    // 1 second
  const m = s * 60  // 1 minute
  const h = m * 60  // 1 hour
  const d = h * 24  // 1 day

  return {s, m, h, d}
}

function parseDuration(duration) {
  const regexp = /^(\d+)([smhd])$/
  const match = regexp.exec(duration)

  const num = match[1]
  const unit = match[2]

  return num * getUnitMap()[unit]
}

function isExpire(saveTime, expire) {
  const expire_ms = parseDuration(expire)
  const time = +new Date
  const r = time - saveTime > expire_ms
  return r
}

export async function localCache (opts) {
  const oldData = getData(opts.key)

  const expire = opts.expire || '24h'

  if ( oldData && !isExpire(oldData.time, expire) ) {
    return oldData.payload
  }

  const payload = await opts.callback()
  const time = +new Date

  setData(
    opts.key,
    {
      time: time,
      payload: payload,
    }
  )

  return payload
}
