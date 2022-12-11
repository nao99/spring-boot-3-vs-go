package org.ndbs.spring.graalvm.product.domain;

import org.ndbs.spring.graalvm.product.api.model.CreateProductDto;
import org.ndbs.spring.graalvm.product.domain.model.Product;
import org.ndbs.spring.graalvm.product.persistence.api.ProductRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Set;

/**
 * ProductServiceImpl class <br>
 *
 * <br>
 *
 * @author Nikolai Osipov <nao99.dev@gmail.com>
 * @since  2022-12-11
 */
@Service
public class ProductServiceImpl implements ProductService {
    private final ProductRepository repository;

    @Autowired
    public ProductServiceImpl(ProductRepository repository) {
        this.repository = repository;
    }

    @Override
    public Set<Product> getProducts() {
        return repository.findLast20();
    }

    @Override
    public Product getProductById(Long productId) {
        return repository.findById(productId)
            .orElseThrow(RuntimeException::new);
    }

    @Override
    public Product createProduct(CreateProductDto createDto) {
        Product product = Product.create(
            createDto.getCategory(),
            createDto.getName(),
            createDto.getDescription()
        );

        return repository.save(product);
    }
}
