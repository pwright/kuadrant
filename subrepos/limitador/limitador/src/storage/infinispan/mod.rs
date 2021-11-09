mod counters;
mod dist_lock;
mod infinispan_storage;
mod response;
mod sets;

use crate::storage::StorageErr;
pub use counters::Consistency;
use infinispan::errors::InfinispanError;
pub use infinispan_storage::InfinispanStorage;
pub use infinispan_storage::InfinispanStorageBuilder;

impl From<reqwest::Error> for StorageErr {
    fn from(e: reqwest::Error) -> Self {
        StorageErr { msg: e.to_string() }
    }
}

impl From<InfinispanError> for StorageErr {
    fn from(e: InfinispanError) -> Self {
        StorageErr { msg: e.to_string() }
    }
}
